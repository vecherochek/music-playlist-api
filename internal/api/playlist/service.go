package playlist

import (
	"github.com/vecherochek/music-playlist-api/internal/service"
	desc "github.com/vecherochek/music-playlist-api/pkg/playlist_v1"
)

type Implementation struct {
	desc.UnimplementedPlaylistV1Server
	playlistService service.PlaylistService
}

func NewImplementation(playlistService service.PlaylistService) *Implementation {
	return &Implementation{
		playlistService: playlistService,
	}
}
