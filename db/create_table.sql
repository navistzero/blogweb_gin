CREATE TABLE IF NOT EXISTS users(
id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
username VARCHAR(64),
password VARCHAR(64),
status INT(4),
createtime INT(10)
);


create table if not exists article(
	id int(4) primary key auto_increment not null,
	title varchar(30),
	author varchar(20),
	tags varchar(30),
	short varchar(255),
	content longtext,
	create_time int(10),
	update_time int(10)
	);