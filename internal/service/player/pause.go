package player

import (
	"context"
	"log"

	"github.com/vecherochek/music-playlist-api/internal/model"
)

func (s *service) Pause(ctx context.Context, playlistUUID string) error {
	player := s.players[playlistUUID]

	if player.Songs.Len() == 0 {
		log.Println("playlist is empty")
		return model.ErrorPlaylistIsEmpty
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
