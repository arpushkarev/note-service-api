package note_v1

import (
	"github.com/arpushkarev/note-service-api/internal/service/note"
	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

// Implementation new server
type Implementation struct {
	desc.UnimplementedNoteV1Server

	noteService *note.Service
}

// NewImplementation starts connect
func NewImplementation(noteService *note.Service) *Implementation {
	return &Implementation{
		noteService: noteService,
	}
}
