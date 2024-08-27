package player

import (
	"context"

	desc "github.com/vecherochek/music-playlist-api/pkg/player_v1"
)

func (i *Implementation) DeletePlayer(ctx context.Context, req *desc.DeletePlayerRequest) (*desc.DeletePlayerResponse, error) {
	err := i.playerService.DeletePlayer(ctx, req.PlayerUuid)
	if err != nil {
		return &desc.DeletePlayerResponse{
			Error: err.Error(),
		}, nil
	}

	return &desc.DeletePlayerResponse{}, nil
}
