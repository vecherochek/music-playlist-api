package song

import (
	"context"

	"github.com/vecherochek/music-playlist-api/internal/converter"
	desc "github.com/vecherochek/music-playlist-api/pkg/playlist_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateSongRequest) (*desc.CreateSongResponse, error) {
	UUID, err := i.songService.Create(ctx, converter.ToSongInfoFromDesc(req.Info))
	if err != nil {
		return nil, err
	}

	return &desc.CreateSongResponse{
		Uuid: UUID,
	}, nil
}
