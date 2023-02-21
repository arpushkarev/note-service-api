package note_v1

import (
	"context"

	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
)

// Note Local structure
type Note struct {
	ID     int64
	Title  string
	Text   string
	Author string
}

// GetAll notes from DataBase
func (i *Implementation) GetAll(ctx context.Context, req *desc.Empty) (*desc.GetAllResponse, error) {

	res, err := i.noteService.GetAll(ctx, req)
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
	//builder := sq.Select("id", "title", "text", "author").
	//	PlaceholderFormat(sq.Dollar).
	//	From(noteTable)
	//
	//query, args, err := builder.ToSql()
	//if err != nil {
	//	return nil, err
	//}
	//
	//row, err := db.QueryContext(ctx, query, args...)
	//if err != nil {
	//	return nil, err
	//}
	//defer row.Close()
	//
	//var res []Note
	//for row.Next() {
	//	var id int64
	//	var (
	//		title, text, author string
	//	)
	//	err = row.Scan(&id, &title, &text, &author)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	res = append(res, Note{
	//		ID:     id,
	//		Title:  title,
	//		Text:   text,
	//		Author: author,
	//	})
	//}
	//
	//var resDesc []*desc.Note
	//for _, elem := range res {
	//	resDesc = append(resDesc, &desc.Note{
	//		Id:     elem.ID,
	//		Title:  elem.Title,
	//		Text:   elem.Text,
	//		Author: elem.Author,
	//	})
	//}
	//
	//return &desc.GetAllResponse{
	//	Notes: resDesc,
	//}, nil
}
