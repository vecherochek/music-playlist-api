package converter

import (
	"encoding/json"

	"github.com/vecherochek/music-playlist-api/internal/model"
	"github.com/vecherochek/music-playlist-api/internal/repository/postgres"
	repoModel "github.com/vecherochek/music-playlist-api/internal/repository/postgres/song/contract/model"
)

func ToSongFromRepo(song *repoModel.Song) (*model.Song, error) {
	var songInfo model.SongInfo
	err := json.Unmarshal(song.SongInfo, &songInfo)
	if err != nil {
		return nil, model.NewError(postgres.ErrorConvertSongFromRepo, err)
	}

	return &model.Song{
		UUID:      song.UUID,
		SongInfo:  songInfo,
		CreatedAt: song.CreatedAt,
		UpdatedAt: song.UpdatedAt,
	}, nil
}

func ToSongFromService(song *model.Song) (*repoModel.Song, error) {
	bytes, err := json.Marshal(song.SongInfo)
	if err != nil {
		return nil, model.NewError(postgres.ErrorConvertSongToRepo, err)
	}

	return &repoModel.Song{
		UUID:      song.UUID,
		SongInfo:  bytes,
		CreatedAt: song.CreatedAt,
		UpdatedAt: song.UpdatedAt,
	}, nil
}
