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
