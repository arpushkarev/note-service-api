package note_v1

import (
	"github.com/arpushkarev/note-service-api/internal/service/note"
	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

// Implementation new server
type Implementation struct {
	desc.UnimplementedNoteV1Server

	NoteService *note.Service
}

// NewImplementation starts connect
func NewImplementation(noteService *note.Service) *Implementation {
	return &Implementation{
		NoteService: noteService,
	}
}

// NewMockNoteV1 initialization
func NewMockNoteV1(i Implementation) *Implementation {
	return &Implementation{
		desc.UnimplementedNoteV1Server{},
		i.NoteService,
	}
}
