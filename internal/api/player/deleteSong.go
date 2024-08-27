package player

import (
	"context"

	desc "github.com/vecherochek/music-playlist-api/pkg/player_v1"
)

func (i *Implementation) DeleteSong(ctx context.Context, req *desc.DeleteSongFromListRequest) (*desc.DeleteSongFromListResponse, error) {
	err := i.playerService.AddSong(ctx, req.PlaylistUuid, req.SongUuid)
	if err != nil {
		return &desc.DeleteSongFromListResponse{
			Error: err.Error(),
		}, nil
	}

	return &desc.DeleteSongFromListResponse{}, nil
}
