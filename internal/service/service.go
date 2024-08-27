package service

import (
	"context"

	"github.com/vecherochek/music-playlist-api/internal/model"
)

type PlaylistService interface {
	AddSong(song *model.Song) error
	Play() error
	Pause() error
	Next() error
	Prev() error
}

type SongService interface {
	Create(ctx context.Context, info *model.SongInfo) (UUID string, err error)
	Get(ctx context.Context, UUID string) (*model.Song, error)
	Update(ctx context.Context, UUID string, info *model.SongInfo) error
	Delete(ctx context.Context, UUID string) error
}
