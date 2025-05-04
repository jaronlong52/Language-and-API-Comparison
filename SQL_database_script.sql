use API_project;

drop table users;
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(20) NOT NULL,
    name VARCHAR(20) NOT NULL
);

insert into users (username, name) values
	("exampleUsername", "exampleName"),
    ("anotherUsername", "anotherName");
    
select * from users;