package model

import (
	"time"
)

type Song struct {
	UUID      string
	SongInfo  []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}
