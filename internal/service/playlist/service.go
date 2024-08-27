package playlist

import (
	"sync"

	"github.com/vecherochek/music-playlist-api/internal/repository"
	def "github.com/vecherochek/music-playlist-api/internal/service"
)

var _ def.PlaylistService = (*service)(nil)

type service struct {
	playlistRepository repository.PlaylistRepository
	m                  sync.RWMutex
}

func NewService(playlistRepository repository.PlaylistRepository) *service {
	return &service{
		playlistRepository: playlistRepository}
}
