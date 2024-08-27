package player

import (
	"context"
	"log"

	"github.com/vecherochek/music-playlist-api/internal/model"
)

func (s *service) Prev(ctx context.Context, playlistUUID string) error {
	s.m.Lock()
	defer s.m.Unlock()

	player, err := s.playerLocalStorage.Get(ctx, playlistUUID)
	if err != nil {
		return err
	}

	if player.Songs.Len() == 0 {
		log.Println("player is empty")
		return model.ErrorPlaylistIsEmpty
	}

	if player.CurrentSong.Prev() == nil {
		log.Println("at the beginning of the list: prev is null")
		return model.ErrorPrevSongNotFound
	}

	if player.Playing {
		err := s.Pause(ctx, playlistUUID)
		if err != nil {
			return err
		}
	}

	player.PrevSong()
	log.Println("switched to the prev song")

	err = s.Play(ctx, playlistUUID)
	if err != nil {
		return err
	}

	return nil
}
