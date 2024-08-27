package model

import "time"

type Song struct {
	UUID     string
	SongInfo SongInfo
}

type SongInfo struct {
	Title    string
	Duration time.Duration
}
