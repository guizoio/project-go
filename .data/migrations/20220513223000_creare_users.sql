-- +goose Up
select 'up SQL query';
create table if not exists users (
    id uuid not null,
    name varchar,
    login varchar,
    password varchar,
    email varchar,
    primary key (id)
);

-- +goose Down
select 'down sql query'
drop table if exists users;