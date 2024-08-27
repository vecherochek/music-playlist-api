package player

import (
	"context"

	desc "github.com/vecherochek/music-playlist-api/pkg/player_v1"
)

func (i *Implementation) Prev(ctx context.Context, req *desc.PrevRequest) (*desc.PrevResponse, error) {

	err := i.playerService.Prev(ctx, req.PlaylistUuid)

	if err != nil {
		return &desc.PrevResponse{
			Error: err.Error(),
		}, nil
	}

	return &desc.PrevResponse{}, nil
}
