package note_v1

import desc "github.com/arpushkarev/note-service-api/pkg/note_v1"

type Note struct {
	desc.UnimplementedNoteV1Server
}

func NewNote() *Note {
	return &Note{}
}

type IdNote struct {
	desc.UnimplementedNoteV1Server
}

func GetNote() *IdNote {
	return &IdNote{}
}

type ListNotes struct {
	desc.UnimplementedNoteV1Server
}

func GetList() *ListNotes {
	return &ListNotes{}
}
