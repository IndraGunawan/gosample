CREATE TABLE IF NOT EXISTS users (
    id int unsigned NOT NULL AUTO_INCREMENT,
    name varchar(100) NOT NULL,
    email varchar(100) NOT NULL,
    password varchar(100) NOT NULL,
    PRIMARY KEY (`id`)
);
