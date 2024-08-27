package player

import (
	"container/list"
	"context"
	"errors"
)

func (s *service) AddPlayer(ctx context.Context, playlistUUID string) (playerUUID string, err error) {
	s.m.Lock()
	defer s.m.Unlock()

	playlist, err := s.playlistRepository.Get(ctx, playlistUUID)
	if err != nil {
		return "", errors.New("Playlist not found")
	}

	var songs list.List

	for _, songUUID := range playlist.PlaylistInfo.SongList {
		song, _ := s.songRepository.Get(ctx, songUUID)
		songs.PushBack(song)
	}

	playerUUID, err = s.playerLocalStorage.Create(ctx, playlistUUID, &songs)
	if err != nil {
		return "", err
	}
	return playerUUID, nil
}
