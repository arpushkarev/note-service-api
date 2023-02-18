-- +goose Up
create table note
(
    id         bigserial primary key,
    title      text,
    text       text,
    author     text,
    created_at timestamp not null default now(),
    updated_at timestamp
);

-- +goose Down
drop table note;
