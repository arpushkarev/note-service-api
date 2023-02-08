package note_v1

import desc "github.com/arpushkarev/note-service-api/pkg/note_v1"

type Implementation struct {
	desc.UnimplementedNoteV1Server
}

func NewImplementation() *Implementation {
	return &Implementation{}
}
