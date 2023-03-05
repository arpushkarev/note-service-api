package note

import (
	"context"

	"github.com/arpushkarev/note-service-api/internal/model"
)

// Update service
func (s *Service) Update(ctx context.Context, updateInfo *model.UpdateNoteInfo) error {
	err := s.noteRepository.Update(ctx, updateInfo)
	if err != nil {
		return err
	}

	return nil
}
