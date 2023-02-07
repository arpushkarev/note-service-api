package main

import (
	"context"
	"log"

	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
	"google.golang.org/grpc"
)

const address = "localhost: 50051"

func main() {
	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("didn't connect: %s", err.Error())
	}
	defer con.Close()

	client := desc.NewNoteV1Client(con)

	resCreate, err := client.Create(context.Background(), &desc.CreateRequest{
		Title:  "Task1",
		Text:   "I've got it",
		Author: "Artem ",
	})
	if err != nil {
		log.Println(err.Error())
	}

	resGet, err := client.Get(context.Background(), &desc.GetRequest{
		Id: 1,
	})
	if err != nil {
		log.Println(err.Error())
	}

	resGetList, err := client.GetList(context.Background(), &desc.GetListRequest{
		ListNotes: "GetList",
	})
	if err != nil {
		log.Println(err.Error())
	}

	resUpdate, err := client.Update(context.Background(), &desc.UpdateRequest{
		Id:     1,
		Title:  "Task1-ruchka4",
		Text:   "Updated",
		Author: "Pushkarev",
	})
	if err != nil {
		log.Println(err.Error())
	}

	resDelete, err := client.Delete(context.Background(), &desc.DeleteRequest{
		Id: 1,
	})
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Id:", resCreate.GetId)
	log.Println("Got title", resGet.GetTitle)
	log.Println("Got text", resGet.GetText)
	log.Println("Got Author", resGet.GetAuthor)
	log.Println("This is the list:", resGetList.GetListId)
	log.Println("Status: ", resUpdate.String())
	log.Println("Status: ", resDelete.String())

}
