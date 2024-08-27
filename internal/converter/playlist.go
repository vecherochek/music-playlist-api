package converter

import (
	"github.com/vecherochek/music-playlist-api/internal/model"
	desc "github.com/vecherochek/music-playlist-api/pkg/player_v1"
)

func ToPlaylistInfoFromDesc(info *desc.PlaylistInfo) *model.PlaylistInfo {
	return &model.PlaylistInfo{
		Name:     info.Name,
		UserId:   info.UserId,
		SongList: info.SongList,
	}
}

func ToPlaylistInfoFromService(info *model.PlaylistInfo) *desc.PlaylistInfo {
	return &desc.PlaylistInfo{
		Name:     info.Name,
		UserId:   info.UserId,
		SongList: info.SongList,
	}
}

func ToPlaylistFromService(song *model.Playlist) *desc.Playlist {
	return &desc.Playlist{
		Uuid:         song.UUID,
		PlaylistInfo: ToPlaylistInfoFromService(&song.PlaylistInfo),
	}
}
