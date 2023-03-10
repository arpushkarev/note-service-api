package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	noteV1 "github.com/arpushkarev/note-service-api/internal/app/api/note_v1"
	"google.golang.org/grpc/credentials/insecure"

	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// App structure
type App struct {
	noteImpl        *noteV1.Implementation
	serviceProvider *serviceProvider
	pathConfig      string
	grpcServer      *grpc.Server
	mux             *runtime.ServeMux
}

// NewApp structure
func NewApp(ctx context.Context, pathConfig string) (*App, error) {
	a := &App{
		pathConfig: pathConfig,
	}

	err := a.initDeps(ctx)

	return a, err
}

// Run servers
func (a *App) Run() error {
	defer func() {
		a.serviceProvider.db.Close()
	}()

	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()

		err := a.startGRPC()
		if err != nil {
			log.Fatalf("GRPCserver error: %s", err.Error())
		}
	}()

	go func() {
		defer wg.Done()

		err := a.startHTTP()
		if err != nil {
			log.Fatalf("HTTPserver error: %s", err.Error())
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initServer,
		a.initGRPCServer,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.pathConfig)

	return nil
}

func (a *App) initServer(ctx context.Context) error {
	a.noteImpl = noteV1.NewImplementation(a.serviceProvider.GetNoteService(ctx))

	return nil
}

func (a *App) initGRPCServer(_ context.Context) error {
	a.grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()),
	)

	desc.RegisterNoteV1Server(a.grpcServer, a.noteImpl)

	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	a.mux = runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := desc.RegisterNoteV1HandlerFromEndpoint(ctx, a.mux, a.serviceProvider.config.GetGRPCAddress(), opts)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) startGRPC() error {
	list, err := net.Listen("tcp", a.serviceProvider.GetConfig().GetGRPCAddress())
	if err != nil {
		return err
	}

	if err = a.grpcServer.Serve(list); err != nil {
		return err
	}

	log.Printf("Started GRPC server on %s host\n", a.serviceProvider.GetConfig().GetGRPCAddress())

	return nil
}

func (a *App) startHTTP() error {
	if err := http.ListenAndServe(a.serviceProvider.GetConfig().GetHTTPAddress(), a.mux); err != nil {
		return err
	}

	log.Printf("Started HTTP server on %s host\n", a.serviceProvider.GetConfig().GetHTTPAddress())

	return nil
}
