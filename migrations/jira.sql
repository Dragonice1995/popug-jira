CREATE SCHEMA IF NOT EXISTS `jira`;

CREATE TABLE IF NOT EXISTS `jira`.`tasks` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `description` VARCHAR(255) NOT NULL,
    `is_ready` TINYINT(1) NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `jira`.`users` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `role` enum('admin', 'manager', 'popug', 'accountant') NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `jira`.`user_tasks` (
    `task_id` INT NOT NULL,
    `user_id` INT NOT NULL,
    CONSTRAINT `task_id_fk` FOREIGN KEY (`task_id`) REFERENCES `tasks`(`id`) ON DELETE CASCADE,
    CONSTRAINT `user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE,
    PRIMARY KEY (`task_id`)
);