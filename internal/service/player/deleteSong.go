package player

import (
	"context"
	"log"

	"github.com/vecherochek/music-playlist-api/internal/model"
)

func (s *service) DeleteSong(ctx context.Context, playlistUUID string, songUUID string) error {
	player, err := s.playerLocalStorage.Get(ctx, playlistUUID)
	if err != nil {
		return err
	}

	player.M.Lock()
	defer player.M.Unlock()

	if player.Playing {
		return model.ErrorDeleteAlreadyPlaying
	}

	err = s.playlistRepository.DeleteSong(ctx, player.PlaylistUUID, songUUID)
	if err != nil {
		return err
	}

	for element := player.Songs.Front(); element != nil; {
		song, ok := element.Value.(*model.Song)
		if ok && song.UUID == songUUID {
			next := element.Next()
			player.Songs.Remove(element)
			element = next
		}
	}

	if player.CurrentSong == nil {
		player.CurrentSong = player.Songs.Front()
	}

	song, err := s.songRepository.Get(ctx, songUUID)
	if err != nil {
		return err
	}

	log.Printf("deleted song: %s\n", song.SongInfo.Title)

	return nil
}
