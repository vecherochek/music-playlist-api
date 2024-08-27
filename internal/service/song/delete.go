package song

import "context"

func (s *service) Delete(ctx context.Context, UUID string) error {
	return s.songRepository.Delete(ctx, UUID)

}
