package player

import (
	"container/list"
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/vecherochek/music-playlist-api/internal/model"
)

func (s *service) AddPlayer(ctx context.Context, playlistUUID string) (playerUUID string, err error) {
	s.m.Lock()
	defer s.m.Unlock()

	playlist, err := s.playlistRepository.Get(ctx, playlistUUID)
	if err != nil {
		return "", errors.New("Playlist not found")
	}

	playerUuid, err := uuid.NewUUID()
	var songs list.List

	for _, songUUID := range playlist.PlaylistInfo.SongList {
		song, _ := s.songRepository.Get(ctx, songUUID)
		songs.PushBack(song)
	}

	player := model.Player{
		PlayerUUID:   playerUuid.String(),
		PlaylistUUID: playlistUUID,
		Songs:        &songs,
		CurrentSong:  songs.Front(),
		Playing:      false,
		PauseChan:    make(chan struct{}),
	}

	s.players[player.PlayerUUID] = &player

	return playerUuid.String(), nil
}
