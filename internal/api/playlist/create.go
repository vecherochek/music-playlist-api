package playlist

import (
	"context"

	"github.com/vecherochek/music-playlist-api/internal/converter"
	desc "github.com/vecherochek/music-playlist-api/pkg/player_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreatePlaylistRequest) (*desc.CreatePlaylistResponse, error) {
	UUID, err := i.playlistService.Create(ctx, converter.ToPlaylistInfoFromDesc(req.Info))
	if err != nil {
		return &desc.CreatePlaylistResponse{
			Error: err.Error(),
		}, nil
	}

	return &desc.CreatePlaylistResponse{
		Uuid: UUID,
	}, nil
}
