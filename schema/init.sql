CREATE DATABASE portfolio;


CREATE TABLE projects (
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL DEFAULT '',
    description VARCHAR(255) NOT NULL DEFAULT '',
    url VARCHAR(255) NOT NULL,
    image VARCHAR(255) NOT NULL DEFAULT ''
);

INSERT INTO projects (title, description, url, image) VALUES ("Test first project", "This is semple project1", "github/1", "../images/project1.jpg");

INSERT INTO projects (title, description, url, image) VALUES ("Test seconf project", "This is semple project2", "github/2", "../images/project2.jpg");

INSERT INTO projects (title, description, url, image) VALUES ("Test three project", "This is semple project3", "github/3", "../images/project3.jpg");

INSERT INTO projects (title, description, url, image) VALUES ("Test four project", "This is semple project4", "github/4", "../images/project4.jpg");

