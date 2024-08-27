package song

import (
	"context"
	"time"

	"github.com/vecherochek/music-playlist-api/internal/model"
)

func (s *service) Create(ctx context.Context, info *model.SongInfo) (UUID string, err error) {
	song := &model.Song{
		SongInfo:  *info,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return s.songRepository.Create(ctx, song)
}
