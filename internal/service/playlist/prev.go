package playlist

import (
	"log"

	"github.com/vecherochek/music-playlist-api/internal/model"
)

func (s *service) Prev() error {
	s.m.Lock()
	defer s.m.Unlock()

	if s.playlist.Songs.Len() == 0 {
		log.Println("playlist is empty")
		return model.ErrorPlaylistIsEmpty
	}

	if s.playlist.CurrentSong.Prev() == nil {
		log.Println("at the beginning of the list: prev is null")
		return model.ErrorPrevSongNotFound
	}

	if s.playlist.Playing {
		err := s.Pause()
		if err != nil {
			return err
		}
	}

	s.playlist.PrevSong()
	log.Println("switched to the prev song")

	err := s.Play()
	if err != nil {
		return err
	}

	return nil
}
