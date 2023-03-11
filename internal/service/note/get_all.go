package note

import (
	"context"

	"github.com/arpushkarev/note-service-api/internal/model"
)

// GetAll service
func (s *Service) GetAll(ctx context.Context) ([]*model.Note, error) {
	res, err := s.noteRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}
