package note

import (
	"context"

	"github.com/arpushkarev/note-service-api/internal/model"
)

// Update service
func (s *Service) Update(ctx context.Context, id int64, updateInfo *model.UpdateNoteInfo) error {
	err := s.noteRepository.Update(ctx, id, updateInfo)
	if err != nil {
		return err
	}

	return nil
}
