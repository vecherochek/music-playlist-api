package playlist

import (
	"context"
	"time"

	"github.com/vecherochek/music-playlist-api/internal/model"
)

func (s *service) Create(ctx context.Context, info *model.PlaylistInfo) (UUID string, err error) {
	playlist := &model.Playlist{
		PlaylistInfo: *info,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return s.playlistRepository.Create(ctx, playlist)
}
