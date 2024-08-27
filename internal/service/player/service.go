package player

import (
	"sync"

	"github.com/vecherochek/music-playlist-api/internal/repository"
	def "github.com/vecherochek/music-playlist-api/internal/service"
)

var _ def.PlayerService = (*service)(nil)

type service struct {
	playerLocalStorage repository.PlayerLocalStorage
	songRepository     repository.SongRepository
	playlistRepository repository.PlaylistRepository
	m                  sync.RWMutex
}

func NewService(songRepository repository.SongRepository, playlistRepository repository.PlaylistRepository, playerLocalStorage repository.PlayerLocalStorage) *service {
	return &service{
		playerLocalStorage: playerLocalStorage,
		songRepository:     songRepository,
		playlistRepository: playlistRepository}
}
