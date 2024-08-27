package app

import (
	"log"

	"github.com/vecherochek/music-playlist-api/internal/api/player"
	"github.com/vecherochek/music-playlist-api/internal/api/playlist"
	"github.com/vecherochek/music-playlist-api/internal/api/song"
	"github.com/vecherochek/music-playlist-api/internal/config"
	"github.com/vecherochek/music-playlist-api/internal/repository"
	"github.com/vecherochek/music-playlist-api/internal/service"
)

type serviceProvider struct {
	grpcConfig config.GRPCConfig

	songRepository repository.SongRepository
	songService    service.SongService
	songImpl       *song.Implementation

	playlistRepository repository.PlaylistRepository
	playlistService    service.PlaylistService
	playlistImpl       *playlist.Implementation

	playerService      service.PlayerService
	playerImpl         *player.Implementation
	playerLocalStorage repository.PlayerLocalStorage
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
