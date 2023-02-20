/*
    CREATE USER blogo_admin@localhost IDENTIFIED BY 'password';
*/

DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS followers;

CREATE TABLE users (
    userId INT auto_increment PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    passwd VARCHAR(128) NOT NULL,
    createdAt TIMESTAMP DEFAULT current_timestamp
) ENGINE=INNODB;

CREATE TABLE posts (
    postId INT auto_increment PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    content VARCHAR(300) NOT NULL,
    authorId INT NOT NULL,
    FOREIGN KEY (authorId) REFERENCES users(userId) ON DELETE CASCADE,
    createdAt TIMESTAMP DEFAULT current_timestamp
) ENGINE=INNODB;

GRANT ALL PRIVILEGES ON Blogo.* TO blogo_admin@localhost;
