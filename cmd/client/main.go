package main

import (
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

	client := desc.NoteV1Client(con)

}
