package main

import (
	"context"
	"log"

	"google.golang.org/grpc/credentials/insecure"

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
		Title:  "Repka",
		Text:   "Posadil ded repku4",
		Author: "Some folk2",
	})
	if err != nil {
		log.Println(err.Error())
	}

	resGet, err := client.Get(ctx, &desc.GetRequest{
		Id: 4,
	})
	if err != nil {
		log.Println(err.Error())
	}

	resGetAll, err := client.GetAll(ctx, &desc.Empty{})
	if err != nil {
		log.Println(err.Error())
	}

	resUpdate, err := client.Update(ctx, &desc.UpdateRequest{
		Id:     4,
		Title:  "Task1-ruchka4",
		Text:   "Updated",
		Author: "Pushkarev",
	})
	if err != nil {
		log.Println(err.Error())
	}

	resDelete, err := client.Delete(ctx, &desc.DeleteRequest{
		Id: 15,
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
