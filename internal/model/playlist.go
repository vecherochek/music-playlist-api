package model

import (
	"container/list"
	"sync"
)

type Playlist struct {
	Songs       *list.List
	CurrentSong *list.Element
	Playing     bool
	PausedAt    float64
	PauseChan   chan struct{}
	m           sync.RWMutex
}

func NewPlaylist() *Playlist {
	return &Playlist{
		Songs:       list.New(),
		CurrentSong: nil,
		Playing:     false,
		PauseChan:   make(chan struct{}),
	}
}

func (playlist *Playlist) NextSong() {
	playlist.m.Lock()
	defer playlist.m.Unlock()

	playlist.CurrentSong = playlist.CurrentSong.Next()
	playlist.PausedAt = 0
}

func (playlist *Playlist) PrevSong() {
	playlist.m.Lock()
	defer playlist.m.Unlock()

	playlist.CurrentSong = playlist.CurrentSong.Prev()
	playlist.PausedAt = 0
}

func (playlist *Playlist) GetCurrentSong() (*Song, error) {
	playlist.m.RLock()
	defer playlist.m.RUnlock()

	return playlist.CurrentSong.Value.(*Song), nil
}
