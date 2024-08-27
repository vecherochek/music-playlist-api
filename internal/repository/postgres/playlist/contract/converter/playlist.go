package converter

import (
	"encoding/json"

	"github.com/vecherochek/music-playlist-api/internal/model"
	"github.com/vecherochek/music-playlist-api/internal/repository/postgres"
	repoModel "github.com/vecherochek/music-playlist-api/internal/repository/postgres/playlist/contract/model"
)

func ToPlaylistFromRepo(playlist *repoModel.Playlist) (*model.Playlist, error) {
	var songInfo model.PlaylistInfo
	err := json.Unmarshal(playlist.PlaylistInfo, &songInfo)
	if err != nil {
		return nil, model.NewError(postgres.ErrorConvertPlaylistFromRepo, err)
	}

	return &model.Playlist{
		UUID:         playlist.UUID,
		PlaylistInfo: songInfo,
		CreatedAt:    playlist.CreatedAt,
		UpdatedAt:    playlist.UpdatedAt,
	}, nil
}

func ToPlaylistFromService(playlist *model.Playlist) (*repoModel.Playlist, error) {
	bytes, err := json.Marshal(playlist.PlaylistInfo)
	if err != nil {
		return nil, model.NewError(postgres.ErrorConvertPlaylistToRepo, err)
	}

	return &repoModel.Playlist{
		UUID:         playlist.UUID,
		PlaylistInfo: bytes,
		CreatedAt:    playlist.CreatedAt,
		UpdatedAt:    playlist.UpdatedAt,
	}, nil
}
