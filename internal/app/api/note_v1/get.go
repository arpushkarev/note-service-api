package note_v1

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
	"github.com/jmoiron/sqlx"
)

//const (
//	noteTable  = "note"
//	host       = "localhost"
//	port       = "54321"
//	dbUser     = "note-service-user"
//	dbPassword = "note-service-password"
//	dbName     = "note-service"
//	sslMode    = "disable"
//)

func (n *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	//fmt.Println("GetNote")
	//fmt.Println("ID:", req.GetId())
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
		From(noteTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()})

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
	var note desc.Note
	err = row.Scan(&note)
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		Note: &desc.Note{
			Id:     note.Id,
			Title:  note.Title,
			Text:   note.Text,
			Author: note.Text,
		},
	}, nil
}
