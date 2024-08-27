package player

import (
	"context"
	"log"

	"github.com/vecherochek/music-playlist-api/internal/model"
)

func (s *service) Pause(ctx context.Context, playlistUUID string) error {
	player, err := s.playerLocalStorage.Get(ctx, playlistUUID)
	if err != nil {
		return err
	}

	if player.Songs.Len() == 0 {
		log.Println("player is empty")
		return model.ErrorPlaylistIsEmpty
	}

	err = player.ReopenLogsChannel()
	if err != nil {
		return nil

	}

	if !player.Playing {
		log.Println("paused already")
		return model.ErrorAlreadyPaused
	}

	player.Playing = false
	player.PauseChan <- struct{}{}

	song, _ := player.GetCurrentSong()

	log.Printf("paused song: %s\n", song.SongInfo.Title)

	return nil
}
