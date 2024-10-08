package song

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/vecherochek/music-playlist-api/internal/model"
	def "github.com/vecherochek/music-playlist-api/internal/repository"
	"github.com/vecherochek/music-playlist-api/internal/repository/postgres/song/contract/converter"
	repoModel "github.com/vecherochek/music-playlist-api/internal/repository/postgres/song/contract/model"
)

var ()

var _ def.SongRepository = (*repository)(nil)

type repository struct {
	db *sql.DB
}

func NewRepository(url string) (*repository, error) {
	db, err := sql.Open("postgres", url)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
		return nil, err
	}

	return &repository{
		db,
	}, nil
}

func (r *repository) Close() {
	err := r.db.Close()

	if err != nil {
		return
	}
}

func (r *repository) Update(ctx context.Context, version time.Time, song *model.Song) error {
	songRepo, err := converter.ToSongFromService(song)
	if err != nil {
		return err
	}

	query, args := UpdateQuery(version, songRepo)
	row := r.db.QueryRowContext(ctx, query, args...)

	err = row.Err()

	if err != nil {
		return model.NewError(def.ErrorUpdateSong, err)
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, songUUID string) error {
	_, err := r.Get(ctx, songUUID)
	if err != nil {
		return model.NewError(def.ErrorGetSong, err)
	}

	query, args := DeleteQuery(songUUID)
	row := r.db.QueryRowContext(ctx, query, args...)

	err = row.Err()
	if err != nil {
		return model.NewError(def.ErrorDeleteSong, err)
	}

	return nil
}

func (r *repository) Create(ctx context.Context, song *model.Song) (UUID string, err error) {
	songRepo, err := converter.ToSongFromService(song)
	if err != nil {
		return "", err
	}

	query, args := CreateQuery(songRepo)
	row := r.db.QueryRowContext(ctx, query, args...)

	var id string
	err = row.Scan(&id)

	if err != nil {
		return "", model.NewError(def.ErrorCreateSong, err)
	}

	return id, nil
}

func (r *repository) Get(ctx context.Context, songUUID string) (*model.Song, error) {
	query, args := GetQuery(songUUID)
	row := r.db.QueryRowContext(ctx, query, args...)

	var song repoModel.Song
	err := row.Scan(&song.UUID, &song.SongInfo, &song.UpdatedAt, &song.CreatedAt)

	if err != nil {
		return nil, model.NewError(def.ErrorGetSong, err)
	}

	return converter.ToSongFromRepo(&song)
}
