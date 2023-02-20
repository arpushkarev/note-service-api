package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/arpushkarev/note-service-api/internal/app/api/note_v1"
	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	hostGrpc = "localhost:50051"
	hostHttp = "localhost:8090"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		startGRPC()
	}()
	go func() {
		defer wg.Done()

		startHTTP()
	}()

	wg.Wait()
}

func startGRPC() error {
	list, err := net.Listen("tcp", hostGrpc)
	if err != nil {
		log.Fatalf("failed to mapping port: %s", err.Error())
	}

	s := grpc.NewServer()
	desc.RegisterNoteV1Server(s, note_v1.NewImplementation())

	fmt.Println("Server is running on port:", hostGrpc)

	if err = s.Serve(list); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
		return err
	}

	return nil
}

func startHTTP() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())} // nolint: staticcheck

	err := desc.RegisterNoteV1HandlerFromEndpoint(ctx, mux, hostGrpc, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(hostHttp, mux)
}
