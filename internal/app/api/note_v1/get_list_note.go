package note_v1

import (
	"context"
	"fmt"
	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

func (n *Note) GetList(ctx context.Context, req *desc.GetListNoteRequest) (*desc.GetListNoteResponse, error) {
	fmt.Println("GetNote")
	fmt.Println("Enter GetList to look all notes:", req.GetListNotes())

	return &desc.GetListNoteResponse{
		ListId: []int64{1, 3, 2},
	}, nil
}
