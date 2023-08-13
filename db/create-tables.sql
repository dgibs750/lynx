USE lynx;

DROP TABLE IF EXISTS account ;
DROP TABLE IF EXISTS user ;

-- -----------------------------------------------------
-- Table lynx.user
-- -----------------------------------------------------

CREATE TABLE IF NOT EXISTS user (
  id INT NOT NULL AUTO_INCREMENT,
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  `master_password` VARCHAR(255) NOT NULL,
  create_time TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `email_UNIQUE` (`email` ASC));


-- -----------------------------------------------------
-- Table lynx.account
-- -----------------------------------------------------

CREATE TABLE IF NOT EXISTS account (
  id INT NOT NULL AUTO_INCREMENT,
  user_id INT NOT NULL,
  account_name VARCHAR(255) NULL,
  username VARCHAR(255) NOT NULL,
  account_password VARCHAR(255) NOT NULL,
  website VARCHAR(255) NULL,
  category VARCHAR(255) NULL,
  create_time TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`, `user_id`),
  INDEX `fk_account_user_idx` (`user_id` ASC),
  CONSTRAINT `fk_account_user`
    FOREIGN KEY (`user_id`)
    REFERENCES user (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);