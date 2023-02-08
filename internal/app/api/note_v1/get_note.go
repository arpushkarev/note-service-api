package note_v1

import (
	"context"
	"fmt"

	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

func (n *Implementation) Get(ctx context.Context, req *desc.Empty) (*desc.GetResponse, error) {
	fmt.Println("GetNote")
	//fmt.Println("ID:", req.GetId())

	return &desc.GetResponse{
		Note: &desc.Note{
			Id:     1,
			Title:  "Funny story",
			Text:   "kolobok povesilsya",
			Author: "Some folk",
		},
	}, nil
}
