-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema journal
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `journal` DEFAULT CHARACTER SET utf8mb3 ;
USE `journal` ;

-- -----------------------------------------------------
-- Table `journal`.`profile`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `journal`.`profile` (
  `idProfile` INT NOT NULL AUTO_INCREMENT,
  `Username` VARCHAR(45) NOT NULL,
  `Password` VARCHAR(45) NOT NULL,
  `Is_Admin` CHAR(1) NOT NULL DEFAULT 'N',
  PRIMARY KEY (`idProfile`),
  UNIQUE INDEX `Profile_UNIQUE` (`Username` ASC) VISIBLE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `journal`.`client`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `journal`.`client` (
  `idProfile` INT NOT NULL,
  `idClient` INT NOT NULL AUTO_INCREMENT,
  `Name` VARCHAR(45) NOT NULL,
  `Description` VARCHAR(100) NOT NULL,
  `StartDate` DATETIME NOT NULL,
  `EndDate` DATETIME NULL DEFAULT NULL,
  PRIMARY KEY (`idClient`),
  UNIQUE INDEX `Client_UNIQUE` (`idProfile` ASC, `idClient` ASC) VISIBLE,
  INDEX `fk_Client_1_idx` (`idProfile` ASC) VISIBLE,
  CONSTRAINT `fk_Client_idProfile`
    FOREIGN KEY (`idProfile`)
    REFERENCES `journal`.`profile` (`idProfile`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `journal`.`project`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `journal`.`project` (
  `idClient` INT NOT NULL,
  `idProject` INT NOT NULL AUTO_INCREMENT,
  `Name` VARCHAR(45) NOT NULL,
  `Description` VARCHAR(100) NOT NULL,
  `StartDate` DATETIME NOT NULL,
  `EndDate` DATETIME NULL DEFAULT NULL,
  `Finished` CHAR(1) NOT NULL DEFAULT 'N',
  PRIMARY KEY (`idProject`),
  UNIQUE INDEX `Project_UNIQUE` (`idClient` ASC, `idProject` ASC) VISIBLE,
  INDEX `fk_idClient_idx` (`idClient` ASC) VISIBLE,
  CONSTRAINT `fk_Project_idClient`
    FOREIGN KEY (`idClient`)
    REFERENCES `journal`.`client` (`idClient`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `journal`.`subproject`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `journal`.`subproject` (
  `idProject` INT NOT NULL,
  `idClient` INT NOT NULL,
  `idSubproject` INT NOT NULL AUTO_INCREMENT,
  `Name` VARCHAR(45) NOT NULL,
  `Description` VARCHAR(100) NOT NULL,
  `StartDate` DATETIME NOT NULL,
  `EndDate` DATETIME NULL DEFAULT NULL,
  `Finished` CHAR(1) NOT NULL DEFAULT 'N',
  PRIMARY KEY (`idSubproject`),
  UNIQUE INDEX `Subproject_UNIQUE` (`idProject` ASC, `idClient` ASC, `idSubproject` ASC) VISIBLE,
  INDEX `fk_Subproject_1_idx` (`idProject` ASC) VISIBLE,
  INDEX `fk_Subproject_2_idx` (`idClient` ASC) VISIBLE,
  CONSTRAINT `fk_Subproject_idClient`
    FOREIGN KEY (`idClient`)
    REFERENCES `journal`.`client` (`idClient`),
  CONSTRAINT `fk_Subproject_idProject`
    FOREIGN KEY (`idProject`)
    REFERENCES `journal`.`project` (`idProject`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `journal`.`journal`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `journal`.`journal` (
  `idProfile` INT NOT NULL,
  `idClient` INT NOT NULL,
  `idProject` INT NOT NULL,
  `idSubproject` INT NOT NULL,
  `idJournal` INT NOT NULL AUTO_INCREMENT,
  `createdAt` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` TIMESTAMP NOT NULL,
  `EntryDate` DATETIME NOT NULL,
  `Description` TEXT NOT NULL,
  `Todos` TEXT NULL DEFAULT NULL,
  `ThingsDone` TEXT NULL DEFAULT NULL,
  `DocUploaded` BLOB NULL DEFAULT NULL,
  PRIMARY KEY (`idJournal`),
  UNIQUE INDEX `Journal_UNIQUE` (`idProfile` ASC, `idClient` ASC, `idProject` ASC, `idSubproject` ASC, `idJournal` ASC, `EntryDate` ASC) VISIBLE,
  INDEX `fk_Journal_2_idx` (`idClient` ASC) VISIBLE,
  INDEX `fk_Journal_3_idx` (`idProject` ASC) VISIBLE,
  INDEX `fk_Journal_4_idx` (`idSubproject` ASC) VISIBLE,
  INDEX `fk_Journal_1_idx` (`idProfile` ASC) VISIBLE,
  CONSTRAINT `fk_Journal_idClient`
    FOREIGN KEY (`idClient`)
    REFERENCES `journal`.`client` (`idClient`),
  CONSTRAINT `fk_Journal_idProfile`
    FOREIGN KEY (`idProfile`)
    REFERENCES `journal`.`profile` (`idProfile`),
  CONSTRAINT `fk_Journal_idProject`
    FOREIGN KEY (`idProject`)
    REFERENCES `journal`.`project` (`idProject`),
  CONSTRAINT `fk_Journal_idSubproject`
    FOREIGN KEY (`idSubproject`)
    REFERENCES `journal`.`subproject` (`idSubproject`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
