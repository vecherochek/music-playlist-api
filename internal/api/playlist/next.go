package playlist

import (
	"context"

	desc "github.com/vecherochek/music-playlist-api/pkg/playlist_v1"
)

func (i *Implementation) Next(ctx context.Context, req *desc.NextRequest) (*desc.NextResponse, error) {

	err := i.playlistService.Next()

	if err != nil {
		return &desc.NextResponse{
			Error: err.Error(),
		}, nil
	}

	return &desc.NextResponse{}, nil
}
