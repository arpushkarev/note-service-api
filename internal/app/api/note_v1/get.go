package note_v1

import (
	"context"

	"github.com/arpushkarev/note-service-api/internal/converter"
	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

// Get note by ID
func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	res, err := i.noteService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		Note: converter.FromModelNote(res),
	}, nil
}
