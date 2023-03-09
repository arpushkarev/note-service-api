package note_v1

import (
	"context"

	"github.com/arpushkarev/note-service-api/internal/converter"
	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Update note by ID
func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	err := i.noteService.Update(
		ctx,
		req.GetId(),
		converter.ToModelUpdateNoteInfo(req.GetNote()),
	)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
