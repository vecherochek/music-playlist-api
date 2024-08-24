package playlist

import (
	"context"

	desc "github.com/vecherochek/music-playlist-api/pkg/playlist_v1"
)

func (i *Implementation) Play(ctx context.Context, req *desc.PlayRequest) (*desc.PlayResponse, error) {

	err := i.playlistService.Play()

	if err != nil {
		return &desc.PlayResponse{
			Error: err.Error(),
		}, nil
	}

	return &desc.PlayResponse{}, nil
}
