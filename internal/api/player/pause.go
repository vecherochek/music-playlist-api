package player

import (
	"context"

	desc "github.com/vecherochek/music-playlist-api/pkg/player_v1"
)

func (i *Implementation) Pause(ctx context.Context, req *desc.PauseRequest) (*desc.PauseResponse, error) {

	err := i.playerService.Pause(ctx, req.PlaylistUuid)

	if err != nil {
		return &desc.PauseResponse{
			Error: err.Error(),
		}, nil
	}

	return &desc.PauseResponse{}, nil
}
