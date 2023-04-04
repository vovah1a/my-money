CREATE TABLE users
(
    id serial primary key,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE board
(
    id serial primary key,
    title varchar(255) not null,
    description varchar(255)
);

CREATE TABLE usersBoard
(
    id serial primary key,
    user_id int references users(id) on delete cascade not null,
    board_id int references board(id) on delete cascade not null
);

CREATE TABLE category
(
    id serial primary key,
    title varchar(255) not null,
    description varchar(255)
);

CREATE TABLE boardsCategory
(
    id serial primary key,
    board_id int references board(id) on delete cascade not null,
    category_id int references category(id) on delete cascade not null
);