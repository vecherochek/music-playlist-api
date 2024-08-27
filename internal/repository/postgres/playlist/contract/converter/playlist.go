package converter

import (
	"encoding/json"

	"github.com/vecherochek/music-playlist-api/internal/model"
	repoModel "github.com/vecherochek/music-playlist-api/internal/repository/postgres/playlist/contract/model"
)

func ToPlaylistFromRepo(playlist *repoModel.Playlist) *model.Playlist {
	var songInfo model.PlaylistInfo
	err := json.Unmarshal(playlist.PlaylistInfo, &songInfo)
	if err != nil {
		return nil
	}

	return &model.Playlist{
		UUID:         playlist.UUID,
		PlaylistInfo: songInfo,
		CreatedAt:    playlist.CreatedAt,
		UpdatedAt:    playlist.UpdatedAt,
	}
}

func ToPlaylistFromService(playlist *model.Playlist) *repoModel.Playlist {
	bytes, _ := json.Marshal(playlist.PlaylistInfo)

	return &repoModel.Playlist{
		UUID:         playlist.UUID,
		PlaylistInfo: bytes,
		CreatedAt:    playlist.CreatedAt,
		UpdatedAt:    playlist.UpdatedAt,
	}
}
