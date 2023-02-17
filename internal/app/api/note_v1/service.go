package note_v1

import desc "github.com/arpushkarev/note-service-api/pkg/note_v1"

// Implementation new server
type Implementation struct {
	desc.UnimplementedNoteV1Server
}

// NewImplementation starts connect
func NewImplementation() *Implementation {
	return &Implementation{}
}
