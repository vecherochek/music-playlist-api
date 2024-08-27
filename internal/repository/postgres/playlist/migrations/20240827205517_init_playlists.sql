-- +goose Up

create table playlists
(
    id           uuid default gen_random_uuid() primary key,
    playlistInfo jsonb null,
    updated_at   timestamp without time zone default now() not null,
    created_at   timestamp without time zone default now() not null
);

create unique index playlists_id_index
    on playlists (id);

-- +goose Down
