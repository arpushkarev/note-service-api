package note

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/arpushkarev/note-service-api/internal/model"
	"github.com/arpushkarev/note-service-api/internal/pkg/db"
	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

const (
	tableName = "note"
)

// Repository - all our handlers
type Repository interface {
	Create(ctx context.Context, noteInfo *model.NoteInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.Note, error)
	GetAll(ctx context.Context, req *desc.Empty) ([]*model.Note, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, req *model.UpdateNoteInfo) error
}

// Repository - db
type repository struct {
	client db.Client
}

// Info structure
type Info struct {
	Title  string
	Text   string
	Author string
}

// Note structure
type Note struct {
	ID   int64
	Info Info
}

// NewRepository - initialisation
func NewRepository(client db.Client) *repository {
	return &repository{
		client: client,
	}
}

// Create new note
func (r *repository) Create(ctx context.Context, noteInfo *model.NoteInfo) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns("title, text, author").
		Values(noteInfo.Title, noteInfo.Text, noteInfo.Author).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "Create",
		QueryRaw: query,
	}

	row, err := r.client.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	row.Next()
	var id int64
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, err
}

// Get the note by ID
func (r *repository) Get(ctx context.Context, id int64) (*model.Note, error) {
	builder := sq.Select("id", "title", "text", "author").
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{"id": id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "Get",
		QueryRaw: query,
	}

	var note model.Note

	err = r.client.DB().GetContext(ctx, &note, q, args...)
	if err != nil {
		return nil, err
	}

	return &note, nil
}

// GetAll notes from DB
func (r *repository) GetAll(ctx context.Context, req *desc.Empty) ([]*model.Note, error) {
	builder := sq.Select("id", "title", "text", "author").
		PlaceholderFormat(sq.Dollar).
		From(tableName)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "GetAll",
		QueryRaw: query,
	}

	var notes []*model.Note

	err = r.client.DB().SelectContext(ctx, &notes, q, args...)
	if err != nil {
		return nil, err
	}

	return notes, nil
}

// Delete the Note by ID
func (r *repository) Delete(ctx context.Context, id int64) error {
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "Update",
		QueryRaw: query,
	}

	res, err := r.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	row := res.RowsAffected()

	if row != 1 {
		log.Printf("expected to affect 1 row, affected %d\n", row)
	}

	return nil
}

// Update the Note by ID
func (r *repository) Update(ctx context.Context, id int64, updateNote *model.UpdateNoteInfo) error {
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id})

	if updateNote.Title.Valid {
		builder = builder.Set("title", updateNote.Title.String)
	}

	if updateNote.Text.Valid {
		builder = builder.Set("text", updateNote.Text.String)
	}

	if updateNote.Author.Valid {
		builder = builder.Set("author", updateNote.Author.String)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "Update",
		QueryRaw: query,
	}

	res, err := r.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	row := res.RowsAffected()

	if row != 1 {
		log.Printf("expected to affect 1 row, affected %d\n", row)
	}

	return nil
}
