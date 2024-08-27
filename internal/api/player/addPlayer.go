package player

import (
	"context"

	desc "github.com/vecherochek/music-playlist-api/pkg/player_v1"
)

func (i *Implementation) AddPlayer(ctx context.Context, req *desc.AddPlayerRequest) (*desc.AddPlayerResponse, error) {
	playerUuid, err := i.playerService.AddPlayer(ctx, req.PlaylistUuid)
	if err != nil {
		return &desc.AddPlayerResponse{
			Error: err.Error(),
		}, nil
	}

	return &desc.AddPlayerResponse{
		PlayerUuid: playerUuid,
	}, nil
}
