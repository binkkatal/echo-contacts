CREATE DATABASE IF NOT EXISTS echo_contacts_dev;
USE echo_contacts_dev;

CREATE TABLE IF NOT EXISTS contacts (
 id INT NOT NULL AUTO_INCREMENT,
 first_name VARCHAR(255) NOT NULL,
 last_name VARCHAR(255) NOT NULL,
 organization VARCHAR(255) NOT NULL,
 phone_number VARCHAR(255) NOT NULL,
 email VARCHAR(255) NOT NULL,
 website VARCHAR(255) NOT NULL,
 PRIMARY KEY (id));
