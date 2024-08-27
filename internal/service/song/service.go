package song

import (
	"github.com/vecherochek/music-playlist-api/internal/repository"
	def "github.com/vecherochek/music-playlist-api/internal/service"
)

var _ def.SongService = (*service)(nil)

type service struct {
	songRepository repository.SongRepository
}

func NewService(
	songRepository repository.SongRepository,
) *service {
	return &service{
		songRepository: songRepository,
	}
}
