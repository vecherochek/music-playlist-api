package player

import (
	"context"
	"log"

	"github.com/vecherochek/music-playlist-api/internal/model"
)

func (s *service) Next(ctx context.Context, playlistUUID string) error {
	s.m.Lock()
	defer s.m.Unlock()

	player := s.players[playlistUUID]

	if player.Songs.Len() == 0 {
		log.Println("playlist is empty")
		return model.ErrorPlaylistIsEmpty
	}

	if player.CurrentSong.Next() == nil {
		log.Println("end of the list: next is null")
		return model.ErrorNextSongNotFound
	}

	if player.Playing {
		err := s.Pause(ctx, playlistUUID)
		if err != nil {
			return err
		}
	}

	player.NextSong()
	log.Println("switched to the next song")

	err := s.Play(ctx, playlistUUID)
	if err != nil {
		return err
	}

	return nil
}
