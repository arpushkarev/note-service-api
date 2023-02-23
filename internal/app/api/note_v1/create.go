package note_v1

import (
	"context"

	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

// Create note with 3 fields Title, Text, Author
func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	res, err := i.noteService.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
