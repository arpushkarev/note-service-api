package note

import (
	"context"

	"github.com/arpushkarev/note-service-api/internal/repository/note"
	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

// Get service
func (s *Service) Get(ctx context.Context, req *desc.GetRequest) (*note.Note, error) {
	note, err := s.noteRepository.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	return note, nil
}
