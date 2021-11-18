CREATE DATABASE abspref_db;

USE abspref_db;

CREATE TABLE IF NOT EXISTS `mydb`.`users_absdb` (
  `users_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `table1col` VARCHAR(45) NULL,
  PRIMARY KEY (`users_id`))
ENGINE = InnoDB
