package note_v1

import (
	"context"

	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

// Delete note by ID
func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*desc.Empty, error) {
	err := i.noteService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.Empty{}, nil
}
