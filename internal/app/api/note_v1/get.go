package note_v1

import (
	"context"

	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

// Get note by ID
func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {

	res, err := i.noteService.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
