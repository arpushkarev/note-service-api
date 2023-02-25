package note

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	//"github.com/arpushkarev/note-service-api/internal/repository/table"
	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
	"github.com/jmoiron/sqlx"
)

const (
	tableName = "note" // tableName - DB name
)

// Repository - all our handlers
type Repository interface {
	Create(ctx context.Context, req *desc.CreateRequest) (int64, error)
	Get(ctx context.Context, req *desc.GetRequest) (*Note, error)
	GetAll(ctx context.Context, req *desc.Empty) ([]*Note, error)
	Delete(ctx context.Context, req *desc.DeleteRequest) error
	Update(ctx context.Context, req *desc.UpdateRequest) error
}

// Repository - db
type repository struct {
	db *sqlx.DB
}

// Note structure
type Note struct {
	ID     int64
	Title  string
	Text   string
	Author string
}

// NewRepository - initialisation
func NewRepository(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

// Create new note
func (r *repository) Create(ctx context.Context, req *desc.CreateRequest) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns("title, text, author").
		Values(req.GetTitle(), req.GetText(), req.GetAuthor()).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
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
func (r *repository) Get(ctx context.Context, req *desc.GetRequest) (*Note, error) {
	builder := sq.Select("id", "title", "text", "author").
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{"id": req.GetId()}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	row.Next()
	var id int64
	var (
		title, text, author string
	)
	err = row.Scan(&id, &title, &text, &author)
	if err != nil {
		return nil, err
	}

	return &Note{
		ID:     id,
		Title:  title,
		Text:   text,
		Author: author,
	}, nil
}

// GetAll notes from DB
func (r *repository) GetAll(ctx context.Context, req *desc.Empty) ([]*Note, error) {
	builder := sq.Select("id", "title", "text", "author").
		PlaceholderFormat(sq.Dollar).
		From(tableName)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var res []Note
	for row.Next() {
		var id int64
		var (
			title, text, author string
		)
		err = row.Scan(&id, &title, &text, &author)
		if err != nil {
			return nil, err
		}

		res = append(res, Note{
			ID:     id,
			Title:  title,
			Text:   text,
			Author: author,
		})
	}

	var resDesc []*Note
	for _, elem := range res {
		resDesc = append(resDesc, &Note{
			ID:     elem.ID,
			Title:  elem.Title,
			Text:   elem.Text,
			Author: elem.Author,
		})
	}

	return resDesc, nil
}

// Delete the Note by ID
func (r *repository) Delete(ctx context.Context, req *desc.DeleteRequest) error {
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	res, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	row, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if row != 1 {
		log.Printf("expected to affect 1 row, affected %d\n", row)
	}

	return nil
}

// Update the Note by ID
func (r *repository) Update(ctx context.Context, req *desc.UpdateRequest) error {
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set("title", req.GetTitle()).
		Set("text", req.GetText()).
		Set("author", req.GetAuthor()).
		Where(sq.Eq{"id": req.GetId()})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	res, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	row, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if row != 1 {
		log.Printf("expected to affect 1 row, affected %d\n", row)
	}

	return nil
}
