package playlist

import (
	"time"

	"github.com/vecherochek/music-playlist-api/internal/repository/postgres/playlist/contract/model"
)

func CreateQuery(playlist *model.Playlist) (query string, args []interface{}) {
	query = `insert into playlists (
                  playlistInfo, 
                  created_at, 
                  updated_at
                  ) 
             values (
                     $1, $2, $3
                   ) 
             returning id;`

	args = []interface{}{playlist.PlaylistInfo, playlist.CreatedAt, playlist.UpdatedAt}

	return query, args
}

func GetQuery(UUID string) (query string, args []interface{}) {
	query = `
		select *
		from playlists
		where 
		    id = $1
		limit 1;`

	args = []interface{}{UUID}

	return query, args
}

func UpdateQuery(version time.Time, playlist *model.Playlist) (query string, args []interface{}) {
	query = `
		UPDATE playlists
		SET
			playlistInfo = $1, 
            updated_at = $2
		WHERE
			id = $3
			and updated_at = $4;`

	args = []interface{}{playlist.PlaylistInfo, playlist.UpdatedAt, playlist.UUID, version}

	return query, args
}
func DeleteQuery(UUID string) (query string, args []interface{}) {
	query = `
			delete 
			from playlists 
       		where 
       		    id = $1;`

	args = []interface{}{UUID}

	return query, args
}
