package player

import (
	"context"

	desc "github.com/vecherochek/music-playlist-api/pkg/player_v1"
)

func (i *Implementation) Next(req *desc.NextRequest, stream desc.PlayerV1_NextServer) error {
	logs, err := i.playerService.Next(context.Background(), req.PlayerUuid)
	if err != nil {
		if err := stream.Send(&desc.StreamNextResponse{Error: err.Error()}); err != nil {
			return err
		}
		return nil
	}

	for log := range logs {
		if err := stream.Send(&desc.StreamNextResponse{Log: log}); err != nil {
			return err
		}
	}

	return nil
}
