package note_v1

import (
	"context"
	"fmt"

	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

func (n *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	fmt.Println("GetNote")
	fmt.Println("ID:", req.GetId())

	return &desc.GetResponse{
		Title:  "Note 2",
		Text:   "Got Note",
		Author: "Art",
	}, nil
}
