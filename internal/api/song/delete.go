package song

import (
	"context"

	desc "github.com/vecherochek/music-playlist-api/pkg/player_v1"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteSongRequest) (*desc.DeleteSongResponse, error) {
	err := i.songService.Delete(ctx, req.GetUuid())
	if err != nil {
		return nil, err
	}

	return &desc.DeleteSongResponse{}, nil
}
