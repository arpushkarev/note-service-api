package note_v1

import (
	"context"
	"fmt"

	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

func (n *Note) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (*desc.DeleteNoteResponse, error) {
	fmt.Println("DeleteNote")
	fmt.Println("Id:", req.GetId())

	return &desc.DeleteNoteResponse{
		DeleteStatus: "Deleted",
	}, nil
}
