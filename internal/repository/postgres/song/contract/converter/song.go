package converter

import (
	"encoding/json"

	"github.com/vecherochek/music-playlist-api/internal/model"
	repoModel "github.com/vecherochek/music-playlist-api/internal/repository/postgres/song/contract/model"
)

func ToSongFromRepo(song *repoModel.Song) *model.Song {
	var songInfo model.SongInfo
	err := json.Unmarshal(song.SongInfo, &songInfo)
	if err != nil {
		return nil
	}

	return &model.Song{
		UUID:      song.UUID,
		SongInfo:  songInfo,
		CreatedAt: song.CreatedAt,
		UpdatedAt: song.UpdatedAt,
	}
}

func ToSongFromService(song *model.Song) *repoModel.Song {
	bytes, _ := json.Marshal(song.SongInfo)

	return &repoModel.Song{
		UUID:      song.UUID,
		SongInfo:  bytes,
		CreatedAt: song.CreatedAt,
		UpdatedAt: song.UpdatedAt,
	}
}
