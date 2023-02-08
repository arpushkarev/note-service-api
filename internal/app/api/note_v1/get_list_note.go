package note_v1

import (
	"context"
	"fmt"

	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

type Note struct {
	Id     int64
	Title  string
	Text   string
	Author string
}

func (n *Implementation) GetList(ctx context.Context, req *desc.GetListRequest) (*desc.GetListResponse, error) {
	fmt.Println("Enter GetList to look all notes:", req.GetNotes())

	return &desc.GetListResponse{
		List: []*desc.Note{
			{
				Id:     1,
				Title:  "Title 1",
				Text:   "Text 1",
				Author: "Author 1",
			},
			{
				Id:     2,
				Title:  "Title 2",
				Text:   "Text 2",
				Author: "Author 2",
			},
		},
	}, nil
}
