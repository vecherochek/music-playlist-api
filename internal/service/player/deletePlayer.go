package player

import "context"

func (s *service) DeletePlayer(ctx context.Context, playerUUID string) error {
	s.m.Lock()
	defer s.m.Unlock()

	delete(s.players, playerUUID)
	return nil
}
