package player

import "context"

func (s *service) DeletePlayer(ctx context.Context, playerUUID string) error {
	s.m.Lock()
	defer s.m.Unlock()

	err := s.Pause(ctx, playerUUID)
	if err != nil {
		return err
	}

	err = s.playerLocalStorage.Delete(ctx, playerUUID)
	if err != nil {
		return err
	}

	return nil
}
