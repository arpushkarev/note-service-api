package main

import (
	"context"
	"flag"
	"log"

	"github.com/arpushkarev/note-service-api/internal/app"
	_ "github.com/jackc/pgx/stdlib" //just for initialization the driver
)

var pathConfig string

func init() {
	flag.StringVar(&pathConfig, "config", "config/config.json", "Path to configuration")
}

func main() {
	flag.Parse()

	ctx := context.Background()
	a, err := app.NewApp(ctx, "")
	if err != nil {
		log.Fatalf("failed to create app %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app %s", err.Error())
	}
}
