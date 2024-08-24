package app

import (
	"log"

	"github.com/vecherochek/music-playlist-api/internal/api/playlist"
	"github.com/vecherochek/music-playlist-api/internal/config"
	"github.com/vecherochek/music-playlist-api/internal/service"
	playlistService "github.com/vecherochek/music-playlist-api/internal/service/playlist"
)

type serviceProvider struct {
	grpcConfig config.GRPCConfig

	//userRepository repository.UserRepository

	userService service.PlaylistService

	playlistImplementation *playlist.Implementation
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
	if s.userService == nil {
		s.userService = playlistService.NewService()
	}

	return s.userService
}

func (s *serviceProvider) PlaylistImpl() *playlist.Implementation {
	if s.playlistImplementation == nil {
		s.playlistImplementation = playlist.NewImplementation(s.PlaylistService())
	}

	return s.playlistImplementation
}
