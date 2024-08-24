package playlist

import (
	"log"
	"time"

	"github.com/vecherochek/music-playlist-api/internal/model"
)

func (s *service) Play() error {
	if s.playlist.Songs.Len() == 0 {
		log.Println("playlist is empty")
		return model.ErrorPlaylistIsEmpty
	}

	if s.playlist.Playing {
		log.Println("playing already")
		return model.ErrorAlreadyPlaying
	}

	s.playlist.Playing = true
	song, _ := s.playlist.GetCurrentSong()

	go s.playSong(song)

	return nil
}

func (s *service) playSong(song *model.Song) {
	songInfo := song.SongInfo

	log.Printf("satrt playing song: %s\n", songInfo.Title)

	for i := s.playlist.PausedAt; i < songInfo.Duration.Seconds(); i++ {
		select {
		case <-s.playlist.PauseChan:
			s.playlist.PausedAt = i
			return
		case <-time.After(time.Second):
			log.Printf(
				"playing %s: %f/%f seconds\n",
				songInfo.Title,
				i+1,
				songInfo.Duration.Seconds())
		}
	}
	s.playlist.Playing = false

	log.Printf("end playing song: %s\n", songInfo.Title)
	log.Printf("start next")

	err := s.Next()
	if err != nil {
		return
	}
}
