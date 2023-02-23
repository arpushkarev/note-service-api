package note_v1

import (
	"context"

	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

// Note Local structure
type Note struct {
	ID     int64
	Title  string
	Text   string
	Author string
}

// GetAll notes
func (i *Implementation) GetAll(ctx context.Context, req *desc.Empty) (*desc.GetAllResponse, error) {
	res, err := i.noteService.GetAll(ctx, req)
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
