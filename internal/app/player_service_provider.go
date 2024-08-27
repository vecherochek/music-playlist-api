package app

import (
	"github.com/vecherochek/music-playlist-api/internal/api/player"
	"github.com/vecherochek/music-playlist-api/internal/repository"
	playerLocalStorage "github.com/vecherochek/music-playlist-api/internal/repository/localcache/player"
	"github.com/vecherochek/music-playlist-api/internal/service"
	playerService "github.com/vecherochek/music-playlist-api/internal/service/player"
)

func (s *serviceProvider) PlayerService() service.PlayerService {
	if s.playerService == nil {
		s.playerService = playerService.NewService(s.SongRepository(), s.PlaylistRepository(), s.PlayerLocalStorage())
	}

	return s.playerService
}

func (s *serviceProvider) PlayerImpl() *player.Implementation {
	if s.playlistImpl == nil {
		s.playerImpl = player.NewImplementation(s.PlayerService())
	}

	return s.playerImpl
}

func (s *serviceProvider) PlayerLocalStorage() repository.PlayerLocalStorage {
	if s.playerLocalStorage == nil {
		s.playerLocalStorage, _ = playerLocalStorage.NewLocalStorage()
	}

	return s.playerLocalStorage
}
