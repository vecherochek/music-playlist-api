package player

import (
	"context"

	desc "github.com/vecherochek/music-playlist-api/pkg/player_v1"
)

func (i *Implementation) Play(ctx context.Context, req *desc.PlayRequest) (*desc.PlayResponse, error) {

	err := i.playerService.Play(ctx, req.PlaylistUuid)

	if err != nil {
		return &desc.PlayResponse{
			Error: err.Error(),
		}, nil
	}

	return &desc.PlayResponse{}, nil
}
