package playlist

import (
	"errors"
	"log"

	"github.com/vecherochek/music-playlist-api/internal/model"
)

func (s *service) Pause() error {
	if s.playlist == nil {
		log.Println("playlist is empty")
		return errors.New("playlist is empty")
	}

	s.playlist.Playing = false
	s.playlist.PauseChan <- struct{}{}

	song := s.playlist.CurrentSong.Value.(*model.Song)

	log.Printf("зause song: %s\n", song.Title)

	return nil
}
