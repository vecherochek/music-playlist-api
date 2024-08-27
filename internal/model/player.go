package model

import (
	"container/list"
	"sync"
)

type Player struct {
	PlayerUUID   string
	PlaylistUUID string
	Songs        *list.List
	CurrentSong  *list.Element
	Playing      bool
	PausedAt     float64
	PauseChan    chan struct{}
	M            sync.RWMutex
}

func NewPlayer(playlistUUID string) *Player {
	return &Player{
		PlaylistUUID: playlistUUID,
		Songs:        list.New(),
		CurrentSong:  nil,
		Playing:      false,
		PauseChan:    make(chan struct{}),
	}
}

func (playlist *Player) NextSong() {
	playlist.M.Lock()
	defer playlist.M.Unlock()

	playlist.CurrentSong = playlist.CurrentSong.Next()
	playlist.PausedAt = 0
}

func (playlist *Player) PrevSong() {
	playlist.M.Lock()
	defer playlist.M.Unlock()

	playlist.CurrentSong = playlist.CurrentSong.Prev()
	playlist.PausedAt = 0
}

func (playlist *Player) GetCurrentSong() (*Song, error) {
	playlist.M.RLock()
	defer playlist.M.RUnlock()

	return playlist.CurrentSong.Value.(*Song), nil
}
