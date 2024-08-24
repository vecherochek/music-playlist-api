package playlist

import (
	"sync"

	"github.com/vecherochek/music-playlist-api/internal/model"
	def "github.com/vecherochek/music-playlist-api/internal/service"
)

var _ def.PlaylistService = (*service)(nil)

type service struct {
	playlist *model.Playlist
	m        sync.RWMutex
}

func NewService() *service {
	return &service{
		playlist: model.NewPlaylist()}
}
