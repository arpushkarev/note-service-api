package note

import (
	"context"

	"github.com/arpushkarev/note-service-api/internal/repository/note"
	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

// GetAll service
func (s *Service) GetAll(ctx context.Context, req *desc.Empty) ([]*note.Note, error) {
	res, err := s.noteRepository.GetAll(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
