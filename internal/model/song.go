package model

import (
	"time"
)

type Song struct {
	UUID            string
	SongInfo        SongInfo
	CreatedAt       time.Time
	UpdatedAt       time.Time
	SongLogsChannel chan string
}

type SongInfo struct {
	Title    string
	Duration time.Duration
}

func (song *Song) UpdateSongInfo(info *SongInfo) {
	song.UpdatedAt = time.Now()
	song.SongInfo = *info
}
