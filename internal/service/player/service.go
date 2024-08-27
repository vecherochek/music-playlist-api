package player

import (
	"sync"

	"github.com/vecherochek/music-playlist-api/internal/model"
	"github.com/vecherochek/music-playlist-api/internal/repository"
	def "github.com/vecherochek/music-playlist-api/internal/service"
)

var _ def.PlayerService = (*service)(nil)

type service struct {
	players            map[string]*model.Player
	songRepository     repository.SongRepository
	playlistRepository repository.PlaylistRepository
	m                  sync.RWMutex
}

func NewService(songRepository repository.SongRepository, playlistRepository repository.PlaylistRepository) *service {
	return &service{
		players:            make(map[string]*model.Player),
		songRepository:     songRepository,
		playlistRepository: playlistRepository}
}
