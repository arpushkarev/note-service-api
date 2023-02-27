package note

import (
	"context"

	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

// Update service
func (s *Service) Update(ctx context.Context, req *desc.UpdateRequest) error {
	err := s.noteRepository.Update(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
