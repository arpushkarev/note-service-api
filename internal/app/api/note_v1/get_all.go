package note_v1

import (
	"context"

	"github.com/arpushkarev/note-service-api/internal/converter"
	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

// GetAll notes
func (i *Implementation) GetAll(ctx context.Context, req *desc.Empty) (*desc.GetAllResponse, error) {
	res, err := i.noteService.GetAll(ctx, req)
	if err != nil {
		return nil, err
	}

	return &desc.GetAllResponse{
		Notes: converter.FromModelNoteSlice(res),
	}, nil
}
