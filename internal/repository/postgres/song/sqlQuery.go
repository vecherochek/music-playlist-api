package song

import (
	"time"

	"github.com/vecherochek/music-playlist-api/internal/repository/postgres/song/contract/model"
)

func CreateQuery(song *model.Song) (query string, args []interface{}) {
	query = `insert into songs (
                  songInfo, 
                  created_at, 
                  updated_at
                  ) 
             values (
                     $1, $2, $3
                   ) 
             returning id;`

	args = []interface{}{song.SongInfo, song.CreatedAt, song.UpdatedAt}

	return query, args
}

func GetQuery(UUID string) (query string, args []interface{}) {
	query = `
		select *
		from songs
		where 
		    id = $1
		limit 1;`

	args = []interface{}{UUID}

	return query, args
}

func UpdateQuery(version time.Time, song *model.Song) (query string, args []interface{}) {
	query = `
		UPDATE songs
		SET
			songInfo = $1, 
            updated_at = $2
		WHERE
			id = $3
			and updated_at = $4;`

	args = []interface{}{song.SongInfo, song.UpdatedAt, song.UUID, version}

	return query, args
}
func DeleteQuery(UUID string) (query string, args []interface{}) {
	query = `
			delete 
			from songs 
       		where 
       		    id = $1;`

	args = []interface{}{UUID}

	return query, args
}
