package app

import (
	"log"

	"github.com/vecherochek/music-playlist-api/internal/api/playlist"
	"github.com/vecherochek/music-playlist-api/internal/api/song"
	"github.com/vecherochek/music-playlist-api/internal/config"
	"github.com/vecherochek/music-playlist-api/internal/repository"
	songRepository "github.com/vecherochek/music-playlist-api/internal/repository/postgres/song"
	"github.com/vecherochek/music-playlist-api/internal/service"
	playlistService "github.com/vecherochek/music-playlist-api/internal/service/playlist"
	songService "github.com/vecherochek/music-playlist-api/internal/service/song"
)

type serviceProvider struct {
	grpcConfig config.GRPCConfig

	playlistService        service.PlaylistService
	playlistImplementation *playlist.Implementation

	songRepository repository.SongRepository
	songService    service.SongService
	songImpl       *song.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) PlaylistService() service.PlaylistService {
	if s.playlistService == nil {
		s.playlistService = playlistService.NewService()
	}

	return s.playlistService
}

func (s *serviceProvider) PlaylistImpl() *playlist.Implementation {
	if s.playlistImplementation == nil {
		s.playlistImplementation = playlist.NewImplementation(s.PlaylistService())
	}

	return s.playlistImplementation
}

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
