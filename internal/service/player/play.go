package player

import (
	"context"
	"log"
	"time"

	"github.com/vecherochek/music-playlist-api/internal/model"
)

func (s *service) Play(ctx context.Context, playlistUUID string) error {

	player := s.players[playlistUUID]

	if player.Songs.Len() == 0 {
		log.Println("playlist is empty")
		return model.ErrorPlaylistIsEmpty
	}

	if player.Playing {
		log.Println("playing already")
		return model.ErrorAlreadyPlaying
	}

	player.Playing = true
	song, _ := player.GetCurrentSong()

	go s.playSong(ctx, song, player)

	return nil
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
			log.Printf(
				"playing %s: %f/%f seconds\n",
				songInfo.Title,
				i+1,
				songInfo.Duration.Seconds())
		}
	}
	player.Playing = false

	log.Printf("end playing song: %s\n", songInfo.Title)
	log.Printf("start next")

	err := s.Next(ctx, player.PlayerUUID)
	if err != nil {
		return
	}
}
