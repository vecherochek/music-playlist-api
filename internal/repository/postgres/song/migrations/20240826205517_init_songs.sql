-- +goose Up

create table if not exists songs
(
    id         uuid default gen_random_uuid() primary key,
    songInfo   jsonb null,
    updated_at timestamp without time zone default now() not null,
    created_at timestamp without time zone default now() not null
);

create unique index if not exists songs_id_index
    on songs (id);

-- +goose Down
