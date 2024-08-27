package app

import (
	"github.com/vecherochek/music-playlist-api/internal/api/playlist"
	"github.com/vecherochek/music-playlist-api/internal/config"
	"github.com/vecherochek/music-playlist-api/internal/repository"
	playlistRepository "github.com/vecherochek/music-playlist-api/internal/repository/postgres/playlist"
	"github.com/vecherochek/music-playlist-api/internal/service"
	playlistService "github.com/vecherochek/music-playlist-api/internal/service/playlist"
)

func (s *serviceProvider) PlaylistService() service.PlaylistService {
	if s.playlistService == nil {
		s.playlistService = playlistService.NewService(
			s.PlaylistRepository())
	}

	return s.playlistService
}

func (s *serviceProvider) PlaylistImpl() *playlist.Implementation {
	if s.playlistImpl == nil {
		s.playlistImpl = playlist.NewImplementation(
			s.PlaylistService())
	}

	return s.playlistImpl
}

func (s *serviceProvider) PlaylistRepository() repository.PlaylistRepository {
	postgresConfig, _ := config.NewPostgresConfig()

	if s.playlistRepository == nil {
		s.playlistRepository, _ = playlistRepository.NewRepository(postgresConfig.Url())
	}

	return s.playlistRepository
}
