package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

type Profile struct {
	Username string
	Password string
	Is_Admin string
}

func main() {
	errEnv := godotenv.Load(".env")

	if errEnv != nil {
		log.Fatalf("Error loading .env file")
	}
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		fmt.Println("Please check the .env file, variable DATABASE_URL=mysql://user:password@host:port/database ")
	}
	dbType := strings.Split(dbUrl, "://")[0]
	userPwdRest := strings.Split(dbUrl, "://")[1]
	userPwd := strings.Split(userPwdRest, "@")[0]
	user := strings.Split(userPwd, ":")[0]
	password := strings.Split(userPwd, ":")[1]
	accessDatabase := strings.Split(userPwdRest, "@")[1]
	access := strings.Split(accessDatabase, "/")[0]
	database := strings.Split(accessDatabase, "/")[1]

	cfg := mysql.Config{
		User:   user,
		Passwd: password,
		Net:    "tcp",
		Addr:   access,
		DBName: database,
	}

	var err error
	db, err = sql.Open(dbType, cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	deleteTables()
	countTables()

	profileID, err := addProfile(Profile{
		Username: "john@john.com",
		Password: "secret",
		Is_Admin: "Y",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added profile: %v\n", profileID)
	countTables()
}

func countTables() ([5]int64, error) {

	var count [5]int64

	rows, err := db.Query("select count(*) from profile")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		if err := rows.Scan(&count[0]); err != nil {
			log.Fatal(err)
		}
	}

	rows, err = db.Query("select count(*) from client")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		if err := rows.Scan(&count[1]); err != nil {
			log.Fatal(err)
		}
	}

	rows, err = db.Query("select count(*) from project")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		if err := rows.Scan(&count[2]); err != nil {
			log.Fatal(err)
		}
	}

	rows, err = db.Query("select count(*) from subproject")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		if err := rows.Scan(&count[3]); err != nil {
			log.Fatal(err)
		}
	}
	rows, err = db.Query("select count(*) from journal")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		if err := rows.Scan(&count[4]); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("profile: %d rows, client: %d rows, project: %d rows, subproject: %d rows, journal: %d rows, \n", count[0], count[1], count[2], count[3], count[4])
	return count, nil
}

func addProfile(profile Profile) (int64, error) {
	result, err := db.Exec("INSERT INTO profile (Username, Password, Is_Admin) VALUES (?, ?, ?)", profile.Username, profile.Password, profile.Is_Admin)
	if err != nil {
		return 0, fmt.Errorf("addProfile: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addProfile: %v", err)
	}
	return id, nil
}

func deleteTables() {
	_, err := db.Exec("delete from journal")
	if err != nil {
		log.Fatal("journal:", err)
	}
	_, err = db.Exec("delete from subproject")
	if err != nil {
		log.Fatal("subproject:", err)
	}
	_, err = db.Exec("delete from project")
	if err != nil {
		log.Fatal("project:", err)
	}
	_, err = db.Exec("delete from client")
	if err != nil {
		log.Fatal("client:", err)
	}
	_, err = db.Exec("delete from profile")
	if err != nil {
		log.Fatal("profile:", err)
	}
}
