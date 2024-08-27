package repository

import (
	"context"
	"time"

	"github.com/vecherochek/music-playlist-api/internal/model"
)

type SongRepository interface {
	Create(ctx context.Context, song *model.Song) (UUID string, err error)
	Get(ctx context.Context, songUUID string) (*model.Song, error)
	Update(ctx context.Context, version time.Time, song *model.Song) error
	Delete(ctx context.Context, songUUID string) error
}

type PlaylistRepository interface {
	Create(ctx context.Context, playlist *model.Playlist) (UUID string, err error)
	Get(ctx context.Context, playlistUUID string) (*model.Playlist, error)
	Update(ctx context.Context, version time.Time, playlist *model.Playlist) error
	Delete(ctx context.Context, playlistUUID string) error

	AddSong(ctx context.Context, playlistUUID string, songUUID string) error
	DeleteSong(ctx context.Context, playlistUUID string, songUUID string) error
}
