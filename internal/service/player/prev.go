package player

import (
	"context"
	"log"

	"github.com/vecherochek/music-playlist-api/internal/model"
)

func (s *service) Prev(ctx context.Context, playlistUUID string) (chan string, error) {
	player, err := s.playerLocalStorage.Get(ctx, playlistUUID)
	if err != nil {
		return nil, err
	}

	player.M.Lock()
	defer player.M.Unlock()

	if player.Songs.Len() == 0 {
		log.Println("player is empty")
		return nil, model.ErrorPlaylistIsEmpty
	}

	if player.CurrentSong.Prev() == nil {
		log.Println("at the beginning of the list: prev is null")
		return nil, model.ErrorPrevSongNotFound
	}

	if player.Playing {
		err := s.Pause(ctx, playlistUUID)
		if err != nil {
			return nil, err
		}
	}

	player.PrevSong()
	log.Println("switched to the prev song")

	return s.Play(ctx, playlistUUID)
}
