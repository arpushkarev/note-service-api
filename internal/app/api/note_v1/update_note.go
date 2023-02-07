package note_v1

import (
	"context"
	"fmt"

	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

func (n *Note) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.UpdateNoteResponse, error) {
	fmt.Println("UpdateNote")
	fmt.Println("Id:", req.GetId())
	fmt.Println("title:", req.GetTitle())
	fmt.Println("text:", req.GetText())
	fmt.Println("author:", req.GetAuthor())

	return &desc.UpdateNoteResponse{
		UpdateStatus: "Finished",
	}, nil
}
