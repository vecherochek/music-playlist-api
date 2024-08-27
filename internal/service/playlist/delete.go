package playlist

import "context"

func (s *service) Delete(ctx context.Context, UUID string) error {
	return s.playlistRepository.Delete(ctx, UUID)
}
