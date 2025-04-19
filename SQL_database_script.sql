use API_project;

drop table users;
create table users (
	username varchar(20),
    name varchar(20)
);

insert into users (username, name) values
	("exampleUsername", "exampleName"),
    ("anotherUsername", "anotherName");
    
select * from users;