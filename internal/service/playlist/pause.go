package playlist

import (
	"log"

	"github.com/vecherochek/music-playlist-api/internal/model"
)

func (s *service) Pause() error {
	if s.playlist.Songs.Len() == 0 {
		log.Println("playlist is empty")
		return model.ErrorPlaylistIsEmpty
	}

	if !s.playlist.Playing {
		log.Println("paused already")
		return model.ErrorAlreadyPaused
	}

	s.playlist.Playing = false
	s.playlist.PauseChan <- struct{}{}

	song, _ := s.playlist.GetCurrentSong()

	log.Printf("paused song: %s\n", song.SongInfo.Title)

	return nil
}
