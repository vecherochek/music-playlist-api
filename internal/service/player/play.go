package player

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/vecherochek/music-playlist-api/internal/model"
)

func (s *service) Play(ctx context.Context, playlistUUID string) (chan string, error) {
	player, err := s.playerLocalStorage.Get(ctx, playlistUUID)
	if err != nil {
		return nil, err
	}

	if player.Songs.Len() == 0 {
		log.Println("player is empty")
		return nil, model.ErrorPlaylistIsEmpty
	}

	if player.Playing {
		log.Println("playing already")
		return nil, model.ErrorAlreadyPlaying
	}

	player.Playing = true
	song, _ := player.GetCurrentSong()

	go s.playSong(ctx, song, player)

	return player.SongLogsChan, nil
}

func (s *service) playSong(ctx context.Context, song *model.Song, player *model.Player) {
	songInfo := song.SongInfo

	log.Printf("satrt playing song: %s\n", songInfo.Title)

	for i := player.PausedAt; i < songInfo.Duration.Seconds(); i++ {
		select {
		case <-player.PauseChan:
			player.PausedAt = i
			return
		case <-time.After(time.Second):
			message := fmt.Sprintf(
				"playing %s: %f/%f seconds\n",
				songInfo.Title,
				i+1,
				songInfo.Duration.Seconds())
			log.Printf(message)
			player.SongLogsChan <- message
		}
	}
	player.Playing = false

	log.Printf("end playing song: %s\n", songInfo.Title)
	log.Printf("start next")

	_, err := s.Next(ctx, player.PlayerUUID)
	if err != nil {
		return
	}
}
