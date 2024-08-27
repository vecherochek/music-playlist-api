package player

import (
	"context"
	"log"

	"github.com/vecherochek/music-playlist-api/internal/model"
)

func (s *service) DeleteSong(ctx context.Context, playlistUUID string, songUUID string) error {
	s.m.Lock()
	defer s.m.Unlock()

	player, err := s.playerLocalStorage.Get(ctx, playlistUUID)
	if err != nil {
		return err
	}

	if player.Playing {
		return model.ErrorDeleteAlreadyPlaying
	}

	song, err := s.songRepository.Get(ctx, songUUID)
	if err != nil {
		return err
	}

	player.Songs.PushBack(song)

	err = s.playlistRepository.DeleteSong(ctx, player.PlaylistUUID, songUUID)
	if err != nil {
		return err
	}

	if player.CurrentSong == nil {
		player.CurrentSong = player.Songs.Front()
	}

	log.Printf("added song: %s\n", song.SongInfo.Title)

	return nil
}
