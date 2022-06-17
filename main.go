package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Profile struct {
	Username string
	Password string
	Is_Admin string
}

type Client struct {
	idProfile   int64
	Name        string
	Description string
	StartDate   string
}

type Project struct {
	idClient    int64
	Name        string
	Description string
	isDefault   string
	StartDate   string
	Finished    string
}
type Subproject struct {
	idProject   int64
	idClient    int64
	Name        string
	Description string
	isDefault   string
	StartDate   string
	Finished    string
}
type Journal struct {
	idProfile    int64
	idClient     int64
	idProject    int64
	idSubproject int64
	EntryDate    string
	Description  string
}

var db *sql.DB
var err error

func main() {
	var errEnv = godotenv.Load(".env")

	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}
	var dbUrl = os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("Please check the .env file, variable DATABASE_URL=mysql://user:password@host:port/database ")
	}
	var dbType = strings.Split(dbUrl, "://")[0]
	var userPwdRest = strings.Split(dbUrl, "://")[1]
	var userPwd = strings.Split(userPwdRest, "@")[0]
	var user = strings.Split(userPwd, ":")[0]
	var password = strings.Split(userPwd, ":")[1]
	var accessDatabase = strings.Split(userPwdRest, "@")[1]
	var access = strings.Split(accessDatabase, "/")[0]
	var database = strings.Split(accessDatabase, "/")[1]

	var cfg = mysql.Config{
		User:   user,
		Passwd: password,
		Net:    "tcp",
		Addr:   access,
		DBName: database,
	}

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
	addProfiles()
	countTables()
}

func addProjects(_idProfile int64, _idClient int64) {
	var projects = []Project{
		Project{
			idClient:    _idClient,
			Name:        "Default project, for client: " + strconv.FormatInt(_idClient, 10),
			Description: "Desc of default project, for client: " + strconv.FormatInt(_idClient, 10),
			isDefault:   "Y",
			StartDate:   time.Now().Format("2006-01-02"),
			Finished:    "N",
		},
		Project{
			idClient:    _idClient,
			Name:        "project 1, for client: " + strconv.FormatInt(_idClient, 10),
			Description: "Description of project 1, for client: " + strconv.FormatInt(_idClient, 10),
			isDefault:   "N",
			StartDate:   time.Now().Format("2006-01-02"),
			Finished:    "N",
		},
	}

	for i := 0; i < len(projects); i++ {
		result, err := db.Exec("INSERT INTO project (idClient,Name, Description, isDefault,StartDate,Finished) VALUES (?,?,?,?,?,?)", projects[i].idClient, projects[i].Name, projects[i].Description, projects[i].isDefault, projects[i].StartDate, projects[i].Finished)
		if err != nil {
			log.Fatalf("addProjects: %v", err)
		}
		idProject, err := result.LastInsertId()
		if err != nil {
			log.Fatalf("addProjects: %v", err)
		}
		//fmt.Printf("ID of added project: %v\n", idProject)
		addSubprojects(_idProfile, idProject, _idClient)
	}
}

func addSubprojects(_idProfile int64, _idProject int64, _idClient int64) {
	var subprojects = []Subproject{
		Subproject{
			idProject:   _idProject,
			idClient:    _idClient,
			Name:        "Default on project: " + strconv.FormatInt(_idProject, 10) + ", for client: " + strconv.FormatInt(_idClient, 10),
			Description: "Description of default subproj, on project: " + strconv.FormatInt(_idProject, 10) + ", for client: " + strconv.FormatInt(_idClient, 10),
			isDefault:   "Y",
			StartDate:   time.Now().Format("2006-01-02"),
			Finished:    "N",
		},
		Subproject{
			idProject:   _idProject,
			idClient:    _idClient,
			Name:        "subproject 1, on project: " + strconv.FormatInt(_idProject, 10) + ", for client: " + strconv.FormatInt(_idClient, 10),
			Description: "Desc of subproj 1, on project: " + strconv.FormatInt(_idProject, 10) + ", for client: " + strconv.FormatInt(_idClient, 10),
			isDefault:   "N",
			StartDate:   time.Now().Format("2006-01-02"),
			Finished:    "N",
		},
	}

	for i := 0; i < len(subprojects); i++ {
		result, err := db.Exec("INSERT INTO subproject (idProject,idClient,Name, Description,isDefault, StartDate,Finished) VALUES (?,?,?,?,?,?,?)", subprojects[i].idProject, subprojects[i].idClient, subprojects[i].Name, subprojects[i].Description, subprojects[i].isDefault, subprojects[i].StartDate, subprojects[i].Finished)
		if err != nil {
			log.Fatalf("addSubprojects: %v", err)
		}
		idSubproject, err := result.LastInsertId()
		if err != nil {
			log.Fatalf("addSubprojects: %v", err)
		}
		//fmt.Printf("ID of added subproject: %v\n", idSubproject)
		addJournals(_idProfile, _idClient, _idProject, idSubproject)
	}
}

func addJournals(_idProfile int64, _idClient int64, _idProject int64, _idSubproject int64) {
	var journals = []Journal{
		Journal{
			idProfile:    _idProfile,
			idClient:     _idClient,
			idProject:    _idProject,
			idSubproject: _idSubproject,
			EntryDate:    time.Now().Format("2006-01-02"),
			Description:  "Journal Entry for " + time.Now().Format("2006-01-02") + " on profile: " + strconv.FormatInt(_idProfile, 10) + ", for client: " + strconv.FormatInt(_idClient, 10) + ", on project: " + strconv.FormatInt(_idProject, 10) + ", subproject: " + strconv.FormatInt(_idSubproject, 10),
		},
		Journal{
			idProfile:    _idProfile,
			idClient:     _idClient,
			idProject:    _idProject,
			idSubproject: _idSubproject,
			EntryDate:    time.Now().Add(time.Hour * 24).Format("2006-01-02"),
			Description:  "Journal Entry for " + time.Now().Add(time.Hour*24).Format("2006-01-02") + " on profile: " + strconv.FormatInt(_idProfile, 10) + ", for client: " + strconv.FormatInt(_idClient, 10) + ", on project: " + strconv.FormatInt(_idProject, 10) + ", subproject: " + strconv.FormatInt(_idSubproject, 10),
		},
	}

	for i := 0; i < len(journals); i++ {
		result, err := db.Exec("INSERT INTO journal (idProfile,idClient,idProject,idSubproject,EntryDate, Description) VALUES (?,?,?,?,?,?)", journals[i].idProfile, journals[i].idClient, journals[i].idProject, journals[i].idSubproject, journals[i].EntryDate, journals[i].Description)
		if err != nil {
			log.Fatalf("addJournals: %v", err)
		}
		_, err = result.LastInsertId()
		if err != nil {
			log.Fatalf("addJournals: %v", err)
		}
		//fmt.Printf("ID of added journal: %v\n", idJournal)
	}
}

func addClients(_idProfile int64) {
	var clients = []Client{
		Client{
			idProfile:   _idProfile,
			Name:        "client 1, on profile " + strconv.FormatInt(_idProfile, 10),
			Description: "Description of client 1, on profile: " + strconv.FormatInt(_idProfile, 10),
			StartDate:   time.Now().Format("2006-01-02"),
		},
		Client{
			idProfile:   _idProfile,
			Name:        "client 2, on profile " + strconv.FormatInt(_idProfile, 10),
			Description: "Description of client 2, on profile: " + strconv.FormatInt(_idProfile, 10),
			StartDate:   time.Now().Format("2006-01-02"),
		},
		Client{
			idProfile:   _idProfile,
			Name:        "client 3, on profile " + strconv.FormatInt(_idProfile, 10),
			Description: "Description of client 3, on profile: " + strconv.FormatInt(_idProfile, 10),
			StartDate:   time.Now().Format("2006-01-02"),
		},
	}

	for i := 0; i < len(clients); i++ {
		result, err := db.Exec("INSERT INTO client (idProfile,Name, Description, StartDate) VALUES (?,?,?,?)", clients[i].idProfile, clients[i].Name, clients[i].Description, clients[i].StartDate)
		if err != nil {
			log.Fatalf("addClients: %v", err)
		}
		idClient, err := result.LastInsertId()
		if err != nil {
			log.Fatalf("addClients: %v", err)
		}
		//fmt.Printf("ID of added client: %v\n", idClient)
		addProjects(_idProfile, idClient)
	}
}

func addProfiles() {
	var profiles = []Profile{
		Profile{
			Username: "marius@gmail.com",
			Password: "secret1",
			Is_Admin: "Y",
		},
		Profile{
			Username: "demo@gmail.com",
			Password: "demo",
			Is_Admin: "N",
		},
		Profile{
			Username: "demo1@gmail.com",
			Password: "demo1",
			Is_Admin: "N",
		},
	}

	for i := 0; i < len(profiles); i++ {
		result, err := db.Exec("INSERT INTO profile (Username, Password, Is_Admin) VALUES (?, ?, ?)", profiles[i].Username, profiles[i].Password, profiles[i].Is_Admin)
		if err != nil {
			log.Fatalf("addProfile: %v", err)
		}
		idProfile, err := result.LastInsertId()
		if err != nil {
			log.Fatalf("addProfile: %v", err)
		}
		//fmt.Printf("ID of added profile: %v\n", idProfile)
		addClients(idProfile)
	}
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
