CREATE TABLE users(
    id serial not null unique,
    email varchar(255) not null unique,
    passwordHash varchar(255) not null
);

CREATE TABLE urls
(
    id serial not null unique,
    url varchar(1024) not null,
    alias varchar(255) not null,
    counter int
);

CREATE TABLE user_url
(
    id serial not null unique,
    user_id int references users(id) on delete cascade not null,
    url_id int references urls(id) on delete cascade not null 
);

CREATE TABLE statistic
(
    id serial not null unique,
    url_id int REFERENCES urls(id) on delete cascade not null,
    ip VARCHAR(255) not null,
    device VARCHAR(255) not null,
    last_date VARCHAR NOT null
);