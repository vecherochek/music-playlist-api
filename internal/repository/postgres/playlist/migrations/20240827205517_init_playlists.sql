-- +goose Up

create table if not exists playlists
(
    id           uuid default gen_random_uuid() primary key,
    playlistInfo jsonb null,
    updated_at   timestamp without time zone default now() not null,
    created_at   timestamp without time zone default now() not null
);

create unique index if not exists playlists_id_index
    on playlists (id);

-- +goose Down
