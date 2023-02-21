package note

import (
	"context"

	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

func (s *Service) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	note, err := s.noteRepository.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		Note: &desc.Note{
			Id:     note.ID,
			Title:  note.Title,
			Text:   note.Text,
			Author: note.Author,
		},
	}, nil
}
