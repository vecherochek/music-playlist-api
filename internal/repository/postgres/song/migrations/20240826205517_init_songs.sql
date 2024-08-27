-- +goose Up

create table songs
(
    id         uuid default gen_random_uuid() primary key,
    songInfo   jsonb null,
    updated_at timestamp without time zone default now() not null,
    created_at timestamp without time zone default now() not null
);

create unique index songs_id_index
    on songs (id);

-- +goose Down
