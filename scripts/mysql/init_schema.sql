create database if not exists account;

use account;

CREATE TABLE if not exists accounts
(
    id         varchar(50) PRIMARY KEY,
    account_id varchar(50) NOT NULL unique,
    name       varchar(50) NOT NULL,
    owner      varchar(50) NOT NULL,
    balance    decimal     NOT NULL,
    created_at datetime    NOT NULL DEFAULT (now())
);

insert into accounts (id, account_id, name, owner, balance) value ('1', '0001-1', 'User test', 'customer', 0);
