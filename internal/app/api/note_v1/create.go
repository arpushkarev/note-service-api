package note_v1

import (
	"context"

	"github.com/arpushkarev/note-service-api/internal/converter"
	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

// Create note with 3 fields Title, Text, Author
func (i *Implementation) Create(ctx context.Context, note *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.noteService.Create(ctx, converter.ToModelNoteInfo(note.GetNote()))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
