package converter

import (
	"time"

	"github.com/golang/protobuf/ptypes/duration"
	"github.com/vecherochek/music-playlist-api/internal/model"
	desc "github.com/vecherochek/music-playlist-api/pkg/playlist_v1"
)

func ToSongInfoFromDesc(info *desc.SongInfo) *model.SongInfo {
	return &model.SongInfo{
		Title:    info.Title,
		Duration: time.Duration(info.Duration.Seconds),
	}
}

func ToSongInfoFromService(info *model.SongInfo) *desc.SongInfo {
	return &desc.SongInfo{
		Title:    info.Title,
		Duration: &duration.Duration{Seconds: info.Duration.Nanoseconds()},
	}
}

func ToSongFromService(song *model.Song) *desc.Song {
	return &desc.Song{
		Uuid:     song.UUID,
		SongInfo: ToSongInfoFromService(&song.SongInfo),
	}
}
