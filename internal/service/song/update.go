package song

import (
	"context"

	"github.com/vecherochek/music-playlist-api/internal/model"
)

func (s *service) Update(ctx context.Context, UUID string, info *model.SongInfo) error {
	song, err := s.Get(ctx, UUID)
	if err != nil {
		return err
	}

	version := song.UpdatedAt
	song.UpdateSongInfo(info)

	return s.songRepository.Update(ctx, version, song)
}
