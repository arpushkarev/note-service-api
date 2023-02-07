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

	res, err := client.Create(context.Background(), &desc.CreateRequest{
		Title:  "Task1",
		Text:   "I've got it",
		Author: "Artem ",
	})
	if err != nil {
		log.Println(err.Error())
	}

	res2, err := client.Get(context.Background(), &desc.GetRequest{
		Id: 1,
	})
	if err != nil {
		log.Println(err.Error())
	}

	res3, err := client.GetList(context.Background(), &desc.GetListRequest{
		ListNotes: "GetList",
	})
	if err != nil {
		log.Println(err.Error())
	}

	res4, err := client.Update(context.Background(), &desc.UpdateRequest{
		Id:     1,
		Title:  "Task1-ruchka4",
		Text:   "Updated",
		Author: "Pushkarev",
	})
	if err != nil {
		log.Println(err.Error())
	}

	res5, err := client.Delete(context.Background(), &desc.DeleteRequest{
		Id: 1,
	})
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Id:", res.Id)
	log.Println("Got title", res2.Title)
	log.Println("Got text", res2.Text)
	log.Println("Got Author", res2.Author)
	log.Println("This is the list:", res3.ListId)
	log.Println("Status: ", res4.String())
	log.Println("Status: ", res5.String())

}
