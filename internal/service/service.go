package service

import (
	"context"

	"github.com/vecherochek/music-playlist-api/internal/model"
)

type PlayerService interface {
	AddPlayer(ctx context.Context, playlistUUID string) (playerUUID string, err error)
	DeletePlayer(ctx context.Context, playerUUID string) error

	AddSong(ctx context.Context, playerUUID string, songUUID string) error
	DeleteSong(ctx context.Context, playerUUID string, songUUID string) error

	Play(ctx context.Context, playerUUID string) (chan string, error)
	Pause(ctx context.Context, playerUUID string) error
	Next(ctx context.Context, playerUUID string) (chan string, error)
	Prev(ctx context.Context, playerUUID string) (chan string, error)
}

type PlaylistService interface {
	Create(ctx context.Context, info *model.PlaylistInfo) (playlistUUID string, err error)
	Get(ctx context.Context, playlistUUID string) (*model.Playlist, error)
	Update(ctx context.Context, playlistUUID string, info *model.PlaylistInfo) error
	Delete(ctx context.Context, playlistUUID string) error
}

type SongService interface {
	Create(ctx context.Context, info *model.SongInfo) (UUID string, err error)
	Get(ctx context.Context, UUID string) (*model.Song, error)
	Update(ctx context.Context, UUID string, info *model.SongInfo) error
	Delete(ctx context.Context, UUID string) error
}
