package contract

import "github.com/vecherochek/music-playlist-api/internal/model"

type SearchQuery struct {
	songUUID []string
	info     *model.SongInfo
}
