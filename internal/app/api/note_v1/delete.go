package note_v1

import (
	"context"

	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

// Delete note by ID
func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*desc.Empty, error) {

	res, err := i.noteService.Delete(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil

	//dbDsn := fmt.Sprintf(
	//	"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
	//	host, port, dbUser, dbPassword, dbName, sslMode,
	//)
	//
	//db, err := sqlx.Open("pgx", dbDsn)
	//if err != nil {
	//	return nil, err
	//}
	//defer db.Close()
	//
	//builder := sq.Delete(noteTable).
	//	PlaceholderFormat(sq.Dollar).
	//	Where(sq.Eq{"id": req.GetId()})
	//
	//query, args, err := builder.ToSql()
	//if err != nil {
	//	return nil, err
	//}
	//
	//res, err := db.ExecContext(ctx, query, args...)
	//if err != nil {
	//	return nil, err
	//}
	//
	//row, err := res.RowsAffected()
	//if err != nil {
	//	return nil, err
	//}
	//
	//if row != 1 {
	//	log.Printf("expected to affect 1 row, affected %d\n", row)
	//}
	//
	//return &desc.Empty{}, nil
}
