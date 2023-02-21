package repository

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/arpushkarev/note-service-api/internal/repository/table"
	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
	"github.com/jmoiron/sqlx"
)

type NoteRepository interface {
	Create(ctx context.Context, req *desc.CreateRequest) (int64, error)
	Get(ctx context.Context, req *desc.GetRequest) (*Note, error)
	GetAll(ctx context.Context, req *desc.Empty) ([]*Note, error)
	Delete(ctx context.Context, req *desc.DeleteRequest) error
	Update(ctx context.Context, req *desc.UpdateRequest) error
}

type Repository struct {
	db *sqlx.DB
}

type Note struct {
	ID     int64
	Title  string
	Text   string
	Author string
}

func NewNoteRepository(db *sqlx.DB) NoteRepository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, req *desc.CreateRequest) (int64, error) {
	builder := sq.Insert(table.Note).
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

func (r *Repository) Get(ctx context.Context, req *desc.GetRequest) (*Note, error) {
	builder := sq.Select("id", "title", "text", "author").
		PlaceholderFormat(sq.Dollar).
		From(table.Note).
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

func (r *Repository) GetAll(ctx context.Context, req *desc.Empty) ([]*Note, error) {
	builder := sq.Select("id", "title", "text", "author").
		PlaceholderFormat(sq.Dollar).
		From(table.Note)

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

func (r *Repository) Delete(ctx context.Context, req *desc.DeleteRequest) error {
	builder := sq.Delete(table.Note).
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

func (r *Repository) Update(ctx context.Context, req *desc.UpdateRequest) error {
	builder := sq.Update(table.Note).
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
