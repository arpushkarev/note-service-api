package note_v1

import (
	"context"

	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
	_ "github.com/jackc/pgx/stdlib" //just for initialization the driver
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

// Create note with 3 objects Title, Text, Author
func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {

	res, err := i.noteService.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil

	//

	//
	//	return &desc.CreateResponse{
	//		Id: id,
	//	}, nil
}
