package playlist

import (
	"context"

	desc "github.com/vecherochek/music-playlist-api/pkg/playlist_v1"
)

func (i *Implementation) Pause(ctx context.Context, req *desc.PauseRequest) (*desc.PauseResponse, error) {

	err := i.playlistService.Pause()

	if err != nil {
		return &desc.PauseResponse{
			Error: err.Error(),
		}, nil
	}

	return &desc.PauseResponse{}, nil
}
