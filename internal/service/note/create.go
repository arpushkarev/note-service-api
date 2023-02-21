package note

import (
	"context"

	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

func (s *Service) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := s.noteRepository.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
