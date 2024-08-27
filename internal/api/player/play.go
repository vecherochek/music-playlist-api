package player

import (
	"context"

	desc "github.com/vecherochek/music-playlist-api/pkg/player_v1"
)

func (i *Implementation) Play(req *desc.PlayRequest, stream desc.PlayerV1_PlayServer) error {
	logs, err := i.playerService.Play(context.Background(), req.PlayerUuid)
	if err != nil {
		if err := stream.Send(&desc.StreamPlayResponse{Error: err.Error()}); err != nil {
			return err
		}
		return nil
	}

	for log := range logs {
		if err := stream.Send(&desc.StreamPlayResponse{Log: log}); err != nil {
			return err
		}
	}

	return nil
}
