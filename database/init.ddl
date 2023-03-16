CREATE DATABASE IF NOT EXISTS todo_app DEFAULT COLLATE utf8mb4_general_ci;
USE todo_app;

CREATE TABLE IF NOT EXISTS `todo` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `create_date` DATETIME DEFAULT(CURDATE()),
    `task` VARCHAR(150),
    PRIMARY KEY (`id`)
);