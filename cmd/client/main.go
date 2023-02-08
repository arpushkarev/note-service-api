package main

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
	"google.golang.org/grpc"
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
		Title:  "Task1",
		Text:   "I've got it",
		Author: "Artem ",
	})
	if err != nil {
		log.Println(err.Error())
	}

	resGet, err := client.Get(ctx, &desc.GetRequest{
		Id: 1,
	})
	if err != nil {
		log.Println(err.Error())
	}

	resGetList, err := client.GetList(ctx, &desc.GetListRequest{
		Notes: "GetList",
	})
	if err != nil {
		log.Println(err.Error())
	}

	resUpdate, err := client.Update(ctx, &desc.UpdateRequest{
		Id:     1,
		Title:  "Task1-ruchka4",
		Text:   "Updated",
		Author: "Pushkarev",
	})
	if err != nil {
		log.Println(err.Error())
	}

	resDelete, err := client.Delete(ctx, &desc.DeleteRequest{
		Id: 1,
	})
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Id:", resCreate.GetId())
	log.Println("Got Title", resGet.GetTitle())
	log.Println("Got Text", resGet.GetText())
	log.Println("Got Author", resGet.GetAuthor())
	log.Println("This is the list:", resGetList.GetList())
	log.Println("Status: ", resUpdate.String())
	log.Println("Status: ", resDelete.String())

}
