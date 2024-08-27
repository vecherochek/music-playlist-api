package playlist

import (
	"context"

	"github.com/vecherochek/music-playlist-api/internal/model"
)

func (s *service) Get(ctx context.Context, playlistUUID string) (*model.Playlist, error) {
	return s.playlistRepository.Get(ctx, playlistUUID)
}
