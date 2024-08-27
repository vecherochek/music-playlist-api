package model

import (
	"time"
)

type Playlist struct {
	UUID         string
	PlaylistInfo PlaylistInfo
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type PlaylistInfo struct {
	Name     string
	UserId   string
	SongList []string
}

func (playlist *Playlist) UpdatePlaylistInfo(info *PlaylistInfo) {
	playlist.UpdatedAt = time.Now()
	playlist.PlaylistInfo = *info
}

func (playlist *Playlist) UpdateSongList(songIdArray []string) {
	playlist.UpdatedAt = time.Now()
	playlist.PlaylistInfo.SongList = songIdArray
}
