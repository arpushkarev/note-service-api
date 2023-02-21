package note

import (
	"context"

	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

func (s *Service) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.Empty, error) {
	err := s.noteRepository.Update(ctx, req)
	if err != nil {
		return nil, err
	}

	return &desc.Empty{}, nil
}
