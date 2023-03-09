package main

import (
	"context"
	"log"

	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
	"github.com/brianvoe/gofakeit/v6"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const address = "localhost: 50051"

func main() {
	con, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("didn't connect: %s", err.Error())
	}
	defer con.Close()

	client := desc.NewNoteV1Client(con)
	ctx := context.Background()

	resCreate, err := client.Create(ctx, &desc.CreateRequest{
		Note: &desc.NoteInfo{
			Title:  gofakeit.Word(),
			Text:   gofakeit.SentenceSimple(),
			Author: gofakeit.Name(),
		},
	})
	if err != nil {
		log.Println(err.Error())
	}

	resGet, err := client.Get(ctx, &desc.GetRequest{
		Id: gofakeit.Int64(),
	})
	if err != nil {
		log.Println(err.Error())
	}

	resGetAll, err := client.GetAll(ctx, &emptypb.Empty{})
	if err != nil {
		log.Println(err.Error())
	}

	resUpdate, err := client.Update(ctx, &desc.UpdateRequest{
		Id: gofakeit.Int64(),
		Note: &desc.UpdateNoteInfo{
			Title:  &wrapperspb.StringValue{Value: gofakeit.Word()},
			Text:   &wrapperspb.StringValue{Value: gofakeit.SentenceSimple()},
			Author: &wrapperspb.StringValue{Value: gofakeit.Name()},
		},
	})
	if err != nil {
		log.Println(err.Error())
	}

	resDelete, err := client.Delete(ctx, &desc.DeleteRequest{
		Id: int64(gofakeit.Number(17, 19)),
	})
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Id:", resCreate.GetId())
	log.Println("Note", resGet.GetNote())
	log.Println("Notes", resGetAll.GetNotes())
	log.Println("Updated successfully", resUpdate.String())
	log.Println("Deleted successfully", resDelete.String())

}
