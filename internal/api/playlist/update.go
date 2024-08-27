package playlist

import (
	"context"

	"github.com/vecherochek/music-playlist-api/internal/converter"
	desc "github.com/vecherochek/music-playlist-api/pkg/player_v1"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdatePlaylistRequest) (*desc.UpdatePlaylistResponse, error) {
	err := i.playlistService.Update(ctx, req.GetUuid(), converter.ToPlaylistInfoFromDesc(req.Info))
	if err != nil {
		return nil, err
	}

	return &desc.UpdatePlaylistResponse{}, nil
}
