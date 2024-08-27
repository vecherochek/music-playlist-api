package playlist

import (
	"context"

	desc "github.com/vecherochek/music-playlist-api/pkg/playlist_v1"
)

func (i *Implementation) Prev(ctx context.Context, req *desc.PrevRequest) (*desc.PrevResponse, error) {

	err := i.playlistService.Prev()

	if err != nil {
		return &desc.PrevResponse{
			Error: err.Error(),
		}, nil
	}

	return &desc.PrevResponse{}, nil
}
