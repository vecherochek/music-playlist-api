package playlist

import (
	"log"

	"github.com/vecherochek/music-playlist-api/internal/model"
)

func (s *service) AddSong(song *model.Song) error {
	s.m.Lock()
	defer s.m.Unlock()

	s.playlist.Songs.PushBack(song)

	if s.playlist.CurrentSong == nil {
		s.playlist.CurrentSong = s.playlist.Songs.Front()
	}

	log.Printf("added song: %s\n", song.SongInfo.Title)

	return nil
}
