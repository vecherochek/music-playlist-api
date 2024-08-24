package playlist

import (
	"log"

	"github.com/vecherochek/music-playlist-api/internal/model"
)

func (s *service) Next() error {
	s.m.Lock()
	defer s.m.Unlock()

	if s.playlist.Songs.Len() == 0 {
		log.Println("playlist is empty")
		return model.ErrorPlaylistIsEmpty
	}

	if s.playlist.CurrentSong.Next() == nil {
		log.Println("end of the list: next is null")
		return model.ErrorNextSongNotFound
	}

	if s.playlist.Playing {
		err := s.Pause()
		if err != nil {
			return err
		}
	}

	s.playlist.NextSong()
	log.Println("switched to the next song")

	err := s.Play()
	if err != nil {
		return err
	}

	return nil
}
