package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/arpushkarev/note-service-api/internal/app/api/note_v1"
	noteRepository "github.com/arpushkarev/note-service-api/internal/repository/note"
	"github.com/arpushkarev/note-service-api/internal/service/note"
	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/jackc/pgx/stdlib" //just for initialization the driver
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	hostGRPC = "50051"
	hostHTTP = "8090"
)

const (
	host       = "localhost"
	port       = "54321"
	dbUser     = "note-service-user"
	dbPassword = "note-service-password"
	dbName     = "note-service"
	sslMode    = "disable"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		err := startGRPC()
		if err != nil {
			log.Fatalf("GRPCserver error: %s", err.Error())
		}
	}()

	go func() {
		defer wg.Done()

		err := startHTTP()
		if err != nil {
			log.Fatalf("HTTPserver error: %s", err.Error())
		}
	}()

	wg.Wait()
}

func startGRPC() error {
	list, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, hostGRPC))
	if err != nil {
		return err
	}

	dbDsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return err
	}
	defer db.Close()

	noteRepository := noteRepository.NewRepository(db)
	noteService := note.NewService(noteRepository)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()),
	)

	desc.RegisterNoteV1Server(s, note_v1.NewImplementation(noteService))

	fmt.Println("GRPC server is running on port:", hostGRPC)

	if err = s.Serve(list); err != nil {
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

	err := desc.RegisterNoteV1HandlerFromEndpoint(ctx, mux, fmt.Sprintf("%s:%s", host, hostGRPC), opts)
	if err != nil {
		return err
	}

	fmt.Println("HTTP server is running on port:", hostHTTP)

	return http.ListenAndServe(fmt.Sprintf("%s:%s", host, hostHTTP), mux)
}
