package song

import (
	"context"

	"github.com/vecherochek/music-playlist-api/internal/converter"
	desc "github.com/vecherochek/music-playlist-api/pkg/player_v1"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateSongRequest) (*desc.UpdateSongResponse, error) {
	err := i.songService.Update(ctx, req.GetUuid(), converter.ToSongInfoFromDesc(req.Info))
	if err != nil {
		return &desc.UpdateSongResponse{
			Error: err.Error(),
		}, nil
	}

	return &desc.UpdateSongResponse{}, nil
}
