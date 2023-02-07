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

	res, err := client.CreateNote(context.Background(), &desc.CreateNoteRequest{
		Title:  "Task1",
		Text:   "I've got it",
		Author: "Artem ",
	})
	if err != nil {
		log.Println(err.Error())
	}

	res2, err := client.GetNote(context.Background(), &desc.GetNoteRequest{
		Id: 1,
	})
	if err != nil {
		log.Println(err.Error())
	}

	res3, err := client.GetListNote(context.Background(), &desc.GetListNoteRequest{
		ListNotes: "GetList",
	})
	if err != nil {
		log.Println(err.Error())
	}

	res4, err := client.UpdateNote(context.Background(), &desc.UpdateNoteRequest{
		Id:     1,
		Title:  "Task1-ruchka4",
		Text:   "Updated",
		Author: "Pushkarev",
	})
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Id:", res.Id)
	log.Println("Got title", res2.Title)
	log.Println("Got text", res2.Text)
	log.Println("Got Author", res2.Author)
	log.Println("This is the list:", res3.ListId)
	log.Println("Status: ", res4.UpdateStatus)

}
