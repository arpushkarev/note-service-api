package note_v1

import (
	"context"

	"github.com/arpushkarev/note-service-api/internal/converter"
	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// GetAll notes
func (i *Implementation) GetAll(ctx context.Context, _ *emptypb.Empty) (*desc.GetAllResponse, error) {
	res, err := i.NoteService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return &desc.GetAllResponse{
		Notes: converter.FromModelNoteSlice(res),
	}, nil
}
