CREATE DATABASE todo;

DROP TABLE IF EXISTS task;
CREATE TABLE task (
    id			serial		 primary key,
	title		varchar(200) not null,
    completed   bit          not null,
    create_ts   timestamp    not null
);

INSERT INTO task(title, completed, create_ts) VALUES('Test Get Task', '1', CURRENT_TIMESTAMP);