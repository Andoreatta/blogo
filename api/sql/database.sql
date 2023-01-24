CREATE USER blogo_admin@localhost IDENTIFIED BY 'password';
CREATE DATABASE IF NOT EXISTS Blogo;
USE Blogo;
GRANT ALL PRIVILEGES ON Blogo.* TO blogo_admin@localhost;
