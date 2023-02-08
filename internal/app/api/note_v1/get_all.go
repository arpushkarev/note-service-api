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

func (n *Implementation) GetAll(ctx context.Context, req *desc.Empty) (*desc.GetAllResponse, error) {
	fmt.Println("GetAll")

	return &desc.GetAllResponse{
		Notes: []*desc.Note{
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
