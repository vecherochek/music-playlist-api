package playlist

import (
	"errors"
	"log"
)

func (s *service) Prev() error {
	if s.playlist == nil {
		log.Println("playlist is empty")
		return errors.New("playlist is empty")
	}

	if s.playlist.CurrentSong.Prev() == nil {
		log.Println("at the beginning of the list: prev is null")
		return errors.New("at the beginning of the list")
	}

	s.playlist.PrevSong()
	log.Println("switched to the prev song")

	err := s.Pause()
	if err != nil {
		return err
	}

	err = s.Play()
	if err != nil {
		return err
	}

	return nil
}
