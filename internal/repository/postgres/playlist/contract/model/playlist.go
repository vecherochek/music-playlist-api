package model

import (
	"time"
)

type Playlist struct {
	UUID         string
	PlaylistInfo []byte
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
