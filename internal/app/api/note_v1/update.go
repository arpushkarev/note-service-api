package note_v1

import (
	"context"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
	"github.com/jmoiron/sqlx"
)

// Update note by ID
func (n *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.Empty, error) {
	dbDsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	builder := sq.Update(noteTable).
		PlaceholderFormat(sq.Dollar).
		Set("title", req.GetTitle()).
		Set("text", req.GetText()).
		Set("author", req.GetAuthor()).
		Where(sq.Eq{"id": req.GetId()})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	res, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	row, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if row != 1 {
		log.Printf("expected to affect 1 row, affected %d\n", row)
	}

	return &desc.Empty{}, nil
}
