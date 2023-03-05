package model

import "database/sql"

type NoteInfo struct {
	Title  string `db:"title"`
	Text   string `db:"text"`
	Author string `db:"author"`
}

type UpdateNoteInfo struct {
	Title  sql.NullString `db:"title"`
	Text   sql.NullString `db:"text"`
	Author sql.NullString `db:"author"`
}

type Note struct {
	Id   int64     `db:"id"`
	Info *NoteInfo `db:""`
}
