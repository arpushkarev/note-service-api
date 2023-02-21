package note

import (
	"context"

	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

func (s *Service) Delete(ctx context.Context, req *desc.DeleteRequest) (*desc.Empty, error) {
	err := s.noteRepository.Delete(ctx, req)
	if err != nil {
		return nil, err
	}

	return &desc.Empty{}, nil
}
