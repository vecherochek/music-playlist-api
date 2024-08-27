package app

import (
	"github.com/vecherochek/music-playlist-api/internal/api/song"
	"github.com/vecherochek/music-playlist-api/internal/config"
	"github.com/vecherochek/music-playlist-api/internal/repository"
	songRepository "github.com/vecherochek/music-playlist-api/internal/repository/postgres/song"
	"github.com/vecherochek/music-playlist-api/internal/service"
	songService "github.com/vecherochek/music-playlist-api/internal/service/song"
)

func (s *serviceProvider) SongRepository() repository.SongRepository {
	postgresConfig, _ := config.NewPostgresConfig()

	if s.songRepository == nil {
		s.songRepository, _ = songRepository.NewRepository(postgresConfig.Url())
	}

	return s.songRepository
}

func (s *serviceProvider) SongService() service.SongService {
	if s.songService == nil {
		s.songService = songService.NewService(
			s.SongRepository())
	}

	return s.songService
}

func (s *serviceProvider) SongImpl() *song.Implementation {
	if s.songImpl == nil {
		s.songImpl = song.NewImplementation(s.SongService())
	}

	return s.songImpl
}
