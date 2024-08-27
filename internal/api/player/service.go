package player

import (
	"github.com/vecherochek/music-playlist-api/internal/service"
	desc "github.com/vecherochek/music-playlist-api/pkg/player_v1"
)

type Implementation struct {
	desc.UnimplementedPlayerV1Server
	playerService service.PlayerService
}

func NewImplementation(playerService service.PlayerService) *Implementation {
	return &Implementation{
		playerService: playerService,
	}
}
