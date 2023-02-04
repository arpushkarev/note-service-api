package main

import (
	"context"
	"log"

	desc "github.com/arpushkarev/Note-Service-Api/pkg/note_v1"
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

	log.Println("Id:", res.Id)

}
