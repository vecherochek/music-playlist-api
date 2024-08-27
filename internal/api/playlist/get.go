package playlist

import (
	"context"

	"github.com/vecherochek/music-playlist-api/internal/converter"
	desc "github.com/vecherochek/music-playlist-api/pkg/player_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetPlaylistRequest) (*desc.GetPlaylistResponse, error) {
	playlist, err := i.playlistService.Get(ctx, req.GetUuid())
	if err != nil {
		return nil, err
	}

	return &desc.GetPlaylistResponse{
		Playlist: converter.ToPlaylistFromService(playlist),
	}, nil
}
