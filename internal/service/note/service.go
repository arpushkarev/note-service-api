package note

import (
	"github.com/arpushkarev/note-service-api/internal/repository/note"
)

// Service structure
type Service struct {
	noteRepository note.Repository
}

// NewService initialisation
func NewService(noteRepository note.Repository) *Service {
	return &Service{
		noteRepository: noteRepository,
	}
}

// NewMockNoteService initialization
func NewMockNoteService(deps ...interface{}) *Service {
	is := Service{}

	for _, v := range deps {
		switch s := v.(type) {
		case note.Repository:
			is.noteRepository = s
		}
	}
	return &is
}
