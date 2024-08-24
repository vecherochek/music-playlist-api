package playlist

import (
	"errors"
	"log"
	"time"

	"github.com/vecherochek/music-playlist-api/internal/model"
)

func (s *service) Play() error {
	if s.playlist == nil {
		log.Println("playlist is empty")
		return errors.New("playlist is empty")
	}

	if s.playlist.Playing {
		log.Println("playing already")
		return errors.New("playing already")
	}

	song := s.playlist.CurrentSong.Value.(*model.Song)

	go s.playSong(song)

	return nil
}

func (s *service) playSong(song *model.Song) {
	s.playlist.Playing = true
	log.Printf("satrt playing song: %s\n", song.Title)

	for i := s.playlist.PausedAt; i < song.Duration.Seconds(); i++ {
		select {
		case <-s.playlist.PauseChan:
			s.playlist.PausedAt = i
			return
		case <-time.After(time.Second):
			log.Printf(
				"playing %s: %f/%f seconds\n",
				song.Title,
				i+1,
				song.Duration.Seconds())
		}
	}
	s.playlist.Playing = false

	log.Printf("end playing song: %s\n", song.Title)

	err := s.Next()
	if err != nil {
		return
	}
}
