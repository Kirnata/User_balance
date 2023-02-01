CREATE SCHEMA IF NOT EXISTS balance;

CREATE TABLE IF NOT EXISTS balance.users
(
    id            bigserial PRIMARY KEY,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null,
    balance       decimal(18,2) check (balance >= 0) not null default 0.00
);

CREATE TABLE IF NOT EXISTS balance.history
(
    id  bigserial  not null unique,
    sender_id int check (sender_id > 0) references balance.users(id) not null,
    getter_id int check (getter_id > 0) references balance.users(id) not null,
    value decimal(18,2) check (value >= 0) not null,
    occurred_at timestamptz NOT NULL,
    description text
);
