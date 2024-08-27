package service

import (
	"github.com/vecherochek/music-playlist-api/internal/model"
)

type PlaylistService interface {
	AddSong(song *model.Song) error
	Play() error
	Pause() error
	Next() error
	Prev() error
}
