DROP TABLE IF EXISTS users;

CREATE TABLE users (
	id varchar(32) PRIMARY KEY,
    password VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

DROP TABLE IF EXISTS posts;

CREATE TABLE posts (
	id varchar(32) PRIMARY KEY,
	post_content varchar(32) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    user_id varchar(32) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

