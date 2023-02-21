package note

import (
	"context"

	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

func (s *Service) GetAll(ctx context.Context, req *desc.Empty) (*desc.GetAllResponse, error) {
	res, err := s.noteRepository.GetAll(ctx, req)
	if err != nil {
		return nil, err
	}

	var resDesc []*desc.Note
	for _, elem := range res {
		resDesc = append(resDesc, &desc.Note{
			Id:     elem.ID,
			Title:  elem.Title,
			Text:   elem.Text,
			Author: elem.Author,
		})
	}

	return &desc.GetAllResponse{
		Notes: resDesc,
	}, nil
}