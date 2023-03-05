package converter

import (
	"database/sql"

	"github.com/arpushkarev/note-service-api/internal/model"
	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func ToModelNoteInfo(info *desc.NoteInfo) *model.NoteInfo {
	return &model.NoteInfo{
		Title:  info.GetTitle(),
		Text:   info.GetText(),
		Author: info.GetAuthor(),
	}
}

func FromModelNoteInfo(model *model.NoteInfo) *desc.NoteInfo {
	return &desc.NoteInfo{
		Title:  model.Title,
		Text:   model.Text,
		Author: model.Author,
	}
}

func FromModelNote(note *model.Note) *desc.Note {
	return &desc.Note{
		Id:   note.Id,
		Note: FromModelNoteInfo(note.Info),
	}
}

func FromModelNoteSlice(notes []*model.Note) []*desc.Note {
	var descNotes []*desc.Note

	for _, elem := range notes {
		descNotes = append(descNotes, FromModelNote(elem))
	}

	return descNotes
}

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

func FromModelUpdateNoteInfo(updateInfo *model.UpdateNoteInfo) *desc.UpdateNoteInfo {
	var title, text, author *wrapperspb.StringValue
	if updateInfo.Title.Valid {
		title = &wrapperspb.StringValue{
			Value: updateInfo.Title.String,
		}
	}
	if updateInfo.Text.Valid {
		text = &wrapperspb.StringValue{
			Value: updateInfo.Text.String,
		}
	}
	if updateInfo.Author.Valid {
		author = &wrapperspb.StringValue{
			Value: updateInfo.Author.String,
		}
	}

	return &desc.UpdateNoteInfo{
		Title:  title,
		Text:   text,
		Author: author,
	}
}
