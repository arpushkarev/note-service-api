package note

import (
	"context"

	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

// Create service
func (s *Service) Create(ctx context.Context, req *desc.CreateRequest) (int64, error) {
	id, err := s.noteRepository.Create(ctx, req)
	if err != nil {
		return 0, err
	}

	return id, err
}
