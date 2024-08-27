package song

import (
	"context"

	"github.com/vecherochek/music-playlist-api/internal/model"
)

func (s *service) Get(ctx context.Context, UUID string) (*model.Song, error) {
	return s.songRepository.Get(ctx, UUID)
}
