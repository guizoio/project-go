-- +goose Up
select 'up SQL query';
create table if not exists addresses (
    id uuid not null,
    street varchar,
    district varchar,
    number varchar,
    city varchar,
    state varchar,
    zip_code varchar,
    enabled boolean default true,
    user_id uuid,
    primary key (id),
    FOREIGN KEY(user_id) REFERENCES users (id) match full
);

-- +goose Down
select 'down sql query'
drop table if exists addresses;