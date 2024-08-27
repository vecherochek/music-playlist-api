package player

import (
	"context"

	desc "github.com/vecherochek/music-playlist-api/pkg/player_v1"
)

func (i *Implementation) AddSong(ctx context.Context, req *desc.AddSongRequest) (*desc.AddSongResponse, error) {
	err := i.playerService.AddSong(ctx, req.PlaylistUuid, req.SongUuid)
	if err != nil {
		return &desc.AddSongResponse{
			Error: err.Error(),
		}, nil
	}

	return &desc.AddSongResponse{}, nil
}
