CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    login VARCHAR(255),
    password VARCHAR(255),
    is_admin BOOLEAN
);

CREATE TABLE langs (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    img_link VARCHAR(255),
    short_description VARCHAR(255),
    author VARCHAR(255),
    year VARCHAR(4),
    version VARCHAR(50),
    description TEXT,
    list TEXT
);

CREATE TABLE projects (
    id SERIAL PRIMARY KEY,
	id_user INT REFERENCES users(id),
    creation_time TIMESTAMP,
    status INT
);

CREATE TABLE files (
    id SERIAL PRIMARY KEY,
    id_lang INT REFERENCES langs(id),
    id_project INT REFERENCES projects(id),
    code TEXT
);