package postgres

import "errors"

var (
	ErrorPlayerNotExists = errors.New("player not exists")
	ErrorCantCreateUUID  = errors.New("can't create uuid")

	ErrorConvertPlaylistFromRepo = errors.New("error converting playlist from repo")
	ErrorConvertPlaylistToRepo   = errors.New("error converting playlist to repo")

	ErrorUpdatePlaylist         = errors.New("error updating playlist")
	ErrorDeletePlaylist         = errors.New("error deleting playlist")
	ErrorGetPlaylist            = errors.New("playlist not found")
	ErrorCreatePlaylist         = errors.New("error creating playlist")
	ErrorDeleteSongFromPlaylist = errors.New("cant delete song from playlist")
	ErrorAddSongToPlaylist      = errors.New("cant add song to playlist")
	ErrorFindSongInPlaylist     = errors.New("no such song in playlist")

	ErrorConvertSongFromRepo = errors.New("error converting song from repo")
	ErrorConvertSongToRepo   = errors.New("error converting song to repo")

	ErrorUpdateSong = errors.New("error updating song")
	ErrorDeleteSong = errors.New("error deleting song")
	ErrorGetSong    = errors.New("song not found")
	ErrorCreateSong = errors.New("error creating song")
)
