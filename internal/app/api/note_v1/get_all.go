package note_v1

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
	"github.com/jmoiron/sqlx"
)

type Note struct {
	Id     int64
	Title  string
	Text   string
	Author string
}

func (n *Implementation) GetAll(ctx context.Context, req *desc.Empty) (*desc.GetAllResponse, error) {
	//fmt.Println("GetAll")
	dbDsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	builder := sq.Select("id", "title", "text", "author").
		PlaceholderFormat(sq.Dollar).
		From(noteTable)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := db.QueryContext(ctx, query, args...)
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

	return &desc.GetAllResponse{
		Notes: []*desc.Note{
			{
				Id:     id,
				Title:  title,
				Text:   text,
				Author: author,
			},
			{
				Id:     id,
				Title:  title,
				Text:   text,
				Author: author,
			},
		},
	}, nil
}
