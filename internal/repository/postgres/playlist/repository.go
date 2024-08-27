package playlist

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/vecherochek/music-playlist-api/internal/model"
	def "github.com/vecherochek/music-playlist-api/internal/repository"
	"github.com/vecherochek/music-playlist-api/internal/repository/postgres/playlist/contract/converter"
	repoModel "github.com/vecherochek/music-playlist-api/internal/repository/postgres/playlist/contract/model"
)

var _ def.PlaylistRepository = (*repository)(nil)

type repository struct {
	db *sql.DB
}

func NewRepository(url string) (*repository, error) {
	db, err := sql.Open("postgres", url)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
		return nil, err
	}

	return &repository{
		db,
	}, nil
}

func (r *repository) Close() {
	err := r.db.Close()

	if err != nil {
		return
	}
}

func (r *repository) Update(ctx context.Context, version time.Time, playlist *model.Playlist) error {
	query, args := UpdateQuery(version, converter.ToPlaylistFromService(playlist))
	row := r.db.QueryRowContext(ctx, query, args...)

	err := row.Err()

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, playlistUUID string) error {
	query, args := DeleteQuery(playlistUUID)
	row := r.db.QueryRowContext(ctx, query, args...)

	err := row.Err()

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Create(ctx context.Context, playlist *model.Playlist) (UUID string, err error) {
	query, args := CreateQuery(converter.ToPlaylistFromService(playlist))
	row := r.db.QueryRowContext(ctx, query, args...)

	var id string
	err = row.Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *repository) Get(ctx context.Context, playlistUUID string) (*model.Playlist, error) {
	query, args := GetQuery(playlistUUID)
	row := r.db.QueryRowContext(ctx, query, args...)

	var playlist repoModel.Playlist
	err := row.Scan(&playlist.UUID, &playlist.PlaylistInfo, &playlist.UpdatedAt, &playlist.CreatedAt)

	if err != nil {
		return nil, errors.New("Playlist not found")
	}

	return converter.ToPlaylistFromRepo(&playlist), nil
}

func (r *repository) AddSong(ctx context.Context, playlistUUID string, songUUID string) error {
	playlist, err := r.Get(ctx, playlistUUID)
	if err != nil {
		return err
	}

	playlist.PlaylistInfo.SongList = append(playlist.PlaylistInfo.SongList, songUUID)

	version := playlist.UpdatedAt
	err = r.Update(ctx, version, playlist)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteSong(ctx context.Context, playlistUUID string, songUUID string) error {
	playlist, err := r.Get(ctx, playlistUUID)
	if err != nil {
		return err
	}

	songs := make([]string, 0, len(playlist.PlaylistInfo.SongList))
	for _, song := range playlist.PlaylistInfo.SongList {
		if song != songUUID {
			songs = append(songs, song)
		}
	}

	if len(songs) == len(playlist.PlaylistInfo.SongList) {
		return errors.New("no such song in playlist")
	}

	playlist.UpdateSongList(songs)
	version := playlist.UpdatedAt

	err = r.Update(ctx, version, playlist)
	if err != nil {
		return err
	}

	return nil
}
