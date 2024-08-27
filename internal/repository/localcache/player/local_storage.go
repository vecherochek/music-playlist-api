package player

import (
	"container/list"
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/vecherochek/music-playlist-api/internal/model"
	def "github.com/vecherochek/music-playlist-api/internal/repository"
)

var _ def.PlayerLocalStorage = (*repository)(nil)

type repository struct {
	players map[string]*model.Player
	m       *sync.RWMutex
}

func NewLocalStorage() (*repository, error) {
	return &repository{
		players: make(map[string]*model.Player),
		m:       new(sync.RWMutex),
	}, nil
}

func (r repository) Create(_ context.Context, playlistUUID string, songs *list.List) (UUID string, err error) {
	r.m.Lock()
	defer r.m.Unlock()

	playerUuid, err := uuid.NewUUID()
	if err != nil {
		return "", errors.New("cannot create uuid")
	}

	player := model.Player{
		PlayerUUID:   playerUuid.String(),
		PlaylistUUID: playlistUUID,
		Songs:        songs,
		CurrentSong:  songs.Front(),
		Playing:      false,
		PauseChan:    make(chan struct{}),
	}

	r.players[player.PlayerUUID] = &player

	return player.PlayerUUID, nil
}

func (r repository) Get(ctx context.Context, playerUUID string) (*model.Player, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	player, ok := r.players[playerUUID]
	if !ok {
		return nil, errors.New("player not exists")
	}

	return player, nil
}

func (r repository) Delete(ctx context.Context, playerUUID string) error {
	r.m.Lock()
	defer r.m.Unlock()

	_, ok := r.players[playerUUID]
	if !ok {
		return errors.New("player not exists")
	}

	delete(r.players, playerUUID)

	return nil
}
