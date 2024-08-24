package converter

import (
	"time"

	"github.com/vecherochek/music-playlist-api/internal/model"
	desc "github.com/vecherochek/music-playlist-api/pkg/playlist_v1"
)

func ToSongInfoFromDesc(info *desc.SongInfo) model.SongInfo {
	return model.SongInfo{
		Title:    info.Title,
		Duration: time.Duration(info.Duration.Seconds) * time.Second,
	}
}
