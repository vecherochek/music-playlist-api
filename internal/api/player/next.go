package player

import (
	"context"

	desc "github.com/vecherochek/music-playlist-api/pkg/player_v1"
)

func (i *Implementation) Next(ctx context.Context, req *desc.NextRequest) (*desc.NextResponse, error) {

	err := i.playerService.Next(ctx, req.PlaylistUuid)

	if err != nil {
		return &desc.NextResponse{
			Error: err.Error(),
		}, nil
	}

	return &desc.NextResponse{}, nil
}
