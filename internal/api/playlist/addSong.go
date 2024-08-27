package playlist

import (
	"context"

	"github.com/vecherochek/music-playlist-api/internal/converter"
	"github.com/vecherochek/music-playlist-api/internal/model"
	desc "github.com/vecherochek/music-playlist-api/pkg/playlist_v1"
)

func (i *Implementation) AddSong(ctx context.Context, req *desc.AddSongRequest) (*desc.AddSongResponse, error) {
	i.playlistService.AddSong(&model.Song{
		SongInfo: converter.ToSongInfoFromDesc(req.Info)})

	return &desc.AddSongResponse{
		Uuid: "11",
	}, nil
}
