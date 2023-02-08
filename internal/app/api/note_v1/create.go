package note_v1

import (
	"context"
	"fmt"

	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

func (n *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	fmt.Println("CreateNote")
	fmt.Println("title:", req.GetTitle())
	fmt.Println("text:", req.GetText())
	fmt.Println("author:", req.GetAuthor())

	return &desc.CreateResponse{
		Id: 1,
	}, nil

}
