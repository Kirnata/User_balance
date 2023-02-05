CREATE TABLE IF NOT EXISTS users
(
    id            bigserial PRIMARY KEY,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null,
    balance       decimal(18,2) check (balance >= 0) not null
);

CREATE TABLE IF NOT EXISTS history
(
    id  bigserial  not null unique,
    sender_id int check (sender_id > 0) references users(id) not null,
    getter_id int check (getter_id > 0) references users(id),
    value decimal(18,2) check (value > 0) not null,

    description text
);
