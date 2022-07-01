CREATE SCHEMA IF NOT EXISTS `jira`;

CREATE TABLE IF NOT EXISTS `jira`.`tasks` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `description` VARCHAR(255) NOT NULL,
  `is_ready` TINYINT(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`));