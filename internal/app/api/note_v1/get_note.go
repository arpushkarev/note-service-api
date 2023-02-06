package note_v1

import (
	"context"
	"fmt"
	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

func (n *IdNote) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	fmt.Println("GetNote")
	fmt.Println("ID:", req.GetId())

	return &desc.GetNoteResponse{
		Title:  "Note 2",
		Text:   "Got Note",
		Author: "Art",
	}, nil
}
