package playlist

import (
	"log"

	"github.com/vecherochek/music-playlist-api/internal/model"
)

func (s *service) AddSong(song *model.Song) error {
	s.playlist.Songs.PushBack(song)

	if s.playlist.CurrentSong == nil {
		s.playlist.CurrentSong = s.playlist.Songs.Front()
	}

	log.Printf("added song: %s\n", song.Title)

	return nil
}
