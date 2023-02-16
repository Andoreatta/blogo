/*
    CREATE USER blogo_admin@localhost IDENTIFIED BY 'password';
*/

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS followers;

CREATE TABLE users (
    userId INT auto_increment PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    passwd varchar(128) NOT NULL,
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE followers (
    userId INT NOT NULL,
        FOREIGN KEY (userId) REFERENCES users(userId)
            ON DELETE CASCADE,

    followerId INT NOT NULL,
        FOREIGN KEY (followerId) REFERENCES users(userId)
            ON DELETE CASCADE,

    PRIMARY KEY (userId, followerId)
) ENGINE=INNODB;

-- CREATE TABLE posts (

-- ) ENGINE=INNODB;

GRANT ALL PRIVILEGES ON Blogo.* TO blogo_admin@localhost;
