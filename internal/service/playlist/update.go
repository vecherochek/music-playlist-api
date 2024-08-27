package playlist

import (
	"context"

	"github.com/vecherochek/music-playlist-api/internal/model"
)

func (s *service) Update(ctx context.Context, playlistUUID string, info *model.PlaylistInfo) error {
	playlist, err := s.Get(ctx, playlistUUID)
	if err != nil {
		return err
	}

	version := playlist.UpdatedAt
	playlist.UpdatePlaylistInfo(info)

	return s.playlistRepository.Update(ctx, version, playlist)
}
