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
