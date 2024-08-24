package playlist

import (
	"errors"
	"log"
)

func (s *service) Next() error {
	if s.playlist == nil {
		log.Println("playlist is empty")
		return errors.New("playlist is empty")
	}

	if s.playlist.CurrentSong.Next() == nil {
		log.Println("end of the list: next is null")
		return errors.New("end of the list")
	}

	s.playlist.NextSong()
	log.Println("switched to the next song")

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
