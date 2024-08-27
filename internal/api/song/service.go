package song

import (
	"github.com/vecherochek/music-playlist-api/internal/service"
	desc "github.com/vecherochek/music-playlist-api/pkg/playlist_v1"
)

type Implementation struct {
	desc.UnimplementedSongV1Server
	songService service.SongService
}

func NewImplementation(songService service.SongService) *Implementation {
	return &Implementation{
		songService: songService,
	}
}
