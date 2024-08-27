package player

import (
	"context"

	desc "github.com/vecherochek/music-playlist-api/pkg/player_v1"
)

func (i *Implementation) Prev(req *desc.PrevRequest, stream desc.PlayerV1_PrevServer) error {
	logs, err := i.playerService.Prev(context.Background(), req.PlayerUuid)
	if err != nil {
		if err := stream.Send(&desc.StreamPrevResponse{Error: err.Error()}); err != nil {
			return err
		}
		return nil
	}

	for log := range logs {
		if err := stream.Send(&desc.StreamPrevResponse{Log: log}); err != nil {
			return err
		}
	}

	return nil
}
