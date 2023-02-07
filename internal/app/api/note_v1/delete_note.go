package note_v1

import (
	"context"
	"fmt"

	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

func (n *Note) Delete(ctx context.Context, req *desc.DeleteRequest) (*desc.Empty, error) {
	fmt.Println("DeleteNote")
	fmt.Println("Id:", req.GetId())

	return &desc.Empty{}, nil
}
