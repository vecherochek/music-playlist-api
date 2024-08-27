package song

import (
	"context"

	"github.com/vecherochek/music-playlist-api/internal/converter"
	desc "github.com/vecherochek/music-playlist-api/pkg/playlist_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetSongRequest) (*desc.GetSongResponse, error) {
	song, err := i.songService.Get(ctx, req.GetUuid())
	if err != nil {
		return nil, err
	}

	return &desc.GetSongResponse{
		Song: converter.ToSongFromService(song),
	}, nil
}
