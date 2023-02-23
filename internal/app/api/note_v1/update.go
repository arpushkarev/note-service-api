package note_v1

import (
	"context"

	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

// Update note by ID
func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.Empty, error) {
	err := i.noteService.Update(ctx, req)
	if err != nil {
		return nil, err
	}

	return &desc.Empty{}, nil
}
