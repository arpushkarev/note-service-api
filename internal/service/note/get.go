package note

import (
	"context"

	"github.com/arpushkarev/note-service-api/internal/model"
)

// Get service
func (s *Service) Get(ctx context.Context, id int64) (*model.Note, error) {
	note, err := s.noteRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return note, nil
}
