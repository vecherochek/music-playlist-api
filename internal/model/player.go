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
	SongLogsChan chan string
	M            sync.RWMutex
}

func NewPlayer(playerUUID string, playlistUUID string, songs *list.List) *Player {
	return &Player{
		PlayerUUID:   playerUUID,
		PlaylistUUID: playlistUUID,
		Songs:        songs,
		CurrentSong:  songs.Front(),
		Playing:      false,
		PauseChan:    make(chan struct{}),
		SongLogsChan: make(chan string, 20),
	}
}

func (player *Player) NextSong() {
	player.CurrentSong = player.CurrentSong.Next()
	player.PausedAt = 0
}

func (player *Player) PrevSong() {
	player.CurrentSong = player.CurrentSong.Prev()
	player.PausedAt = 0
}

func (player *Player) GetCurrentSong() (*Song, error) {
	return player.CurrentSong.Value.(*Song), nil
}

func (player *Player) ReopenLogsChannel() error {
	close(player.SongLogsChan)
	player.SongLogsChan = make(chan string, 20)
	return nil
}
