package playlist

import (
	"context"

	desc "github.com/vecherochek/music-playlist-api/pkg/player_v1"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeletePlaylistRequest) (*desc.DeletePlaylistResponse, error) {
	err := i.playlistService.Delete(ctx, req.GetUuid())
	if err != nil {
		return &desc.DeletePlaylistResponse{
			Error: err.Error(),
		}, nil
	}

	return &desc.DeletePlaylistResponse{}, nil
}
