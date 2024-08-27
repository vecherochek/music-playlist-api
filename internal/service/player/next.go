package player

import (
	"context"
	"log"

	"github.com/vecherochek/music-playlist-api/internal/model"
)

func (s *service) Next(ctx context.Context, playlistUUID string) (chan string, error) {
	player, err := s.playerLocalStorage.Get(ctx, playlistUUID)
	if err != nil {
		return nil, err
	}

	player.M.Lock()
	defer player.M.Unlock()

	if player.Songs.Len() == 0 {
		log.Println("playlist is empty")
		return nil, model.ErrorPlaylistIsEmpty
	}

	if player.CurrentSong.Next() == nil {
		log.Println("end of the list: next is null")
		err := s.Pause(ctx, playlistUUID)
		if err != nil {
			return nil, err
		}
		return nil, model.ErrorNextSongNotFound
	}

	if player.Playing {
		err := s.Pause(ctx, playlistUUID)
		if err != nil {
			return nil, err
		}
	}

	player.NextSong()
	player.PausedAt = 0
	log.Println("switched to the next song")

	return s.Play(ctx, playlistUUID)
}
