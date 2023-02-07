package note_v1

import (
	"context"
	"fmt"

	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

func (n *Implementation) GetList(ctx context.Context, req *desc.GetListRequest) (*desc.GetListResponse, error) {
	fmt.Println("Enter GetList to look all notes:", req.GetListNotes())

	return &desc.GetListResponse{
		ListId: []int64{1, 3, 2},
	}, nil
}
