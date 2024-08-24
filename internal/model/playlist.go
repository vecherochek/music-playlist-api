package model

import (
	"container/list"
	"sync"
)

type Playlist struct {
	Songs       *list.List
	CurrentSong *list.Element
	mu          sync.Mutex
	Playing     bool
	PausedAt    float64
	PauseChan   chan struct{}
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
	playlist.CurrentSong = playlist.CurrentSong.Next()
	playlist.PausedAt = 0
}

func (playlist *Playlist) PrevSong() {
	playlist.CurrentSong = playlist.CurrentSong.Prev()
	playlist.PausedAt = 0
}
