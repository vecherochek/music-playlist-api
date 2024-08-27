package model

import "errors"

var (
	ErrorPlaylistIsEmpty      = errors.New("playlist is empty")
	ErrorPrevSongNotFound     = errors.New("prev song not found, current song at the beginning of the playlist")
	ErrorNextSongNotFound     = errors.New("next song not found, current song at the end of the playlist")
	ErrorAlreadyPaused        = errors.New("nothing is being played")
	ErrorAlreadyPlaying       = errors.New("the song is already playing")
	ErrorDeleteAlreadyPlaying = errors.New("cannot delete song because is already playing")
)

func NewError(template error, err error) error {
	return errors.New(template.Error() + err.Error())
}
