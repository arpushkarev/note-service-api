package converter

import (
	"database/sql"

	"github.com/arpushkarev/note-service-api/internal/model"
	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToModelNoteInfo converts structure from client query into model
func ToModelNoteInfo(info *desc.NoteInfo) *model.NoteInfo {
	return &model.NoteInfo{
		Title:  info.GetTitle(),
		Text:   info.GetText(),
		Author: info.GetAuthor(),
	}
}

// FromModelNoteInfo converts structure from model into client response
func FromModelNoteInfo(model *model.NoteInfo) *desc.NoteInfo {
	return &desc.NoteInfo{
		Title:  model.Title,
		Text:   model.Text,
		Author: model.Author,
	}
}

// FromModelNote converts structure from model into client response
func FromModelNote(note *model.Note) *desc.Note {
	var updatedAt *timestamppb.Timestamp
	if note.UpdatedAt.Valid {
		updatedAt = timestamppb.New(note.UpdatedAt.Time)
	}

	return &desc.Note{
		Id:        note.ID,
		Note:      FromModelNoteInfo(note.Info),
		CreatedAt: timestamppb.New(note.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

// FromModelNoteSlice converts slice of structures from model into client response
func FromModelNoteSlice(notes []*model.Note) []*desc.Note {
	var descNotes []*desc.Note

	for _, elem := range notes {
		descNotes = append(descNotes, FromModelNote(elem))
	}

	return descNotes
}

// ToModelUpdateNoteInfo converts structure from client's query into model
func ToModelUpdateNoteInfo(updateInfo *desc.UpdateNoteInfo) *model.UpdateNoteInfo {
	var title, text, author sql.NullString
	if updateInfo.GetTitle() != nil {
		title = sql.NullString{
			String: updateInfo.GetTitle().GetValue(),
			Valid:  updateInfo.GetTitle() != nil,
		}
	}

	if updateInfo.GetText() != nil {
		text = sql.NullString{
			String: updateInfo.GetText().GetValue(),
			Valid:  updateInfo.GetText() != nil,
		}
	}

	if updateInfo.GetAuthor() != nil {
		author = sql.NullString{
			String: updateInfo.GetAuthor().GetValue(),
			Valid:  updateInfo.GetAuthor() != nil,
		}
	}

	return &model.UpdateNoteInfo{
		Title:  title,
		Text:   text,
		Author: author,
	}
}
