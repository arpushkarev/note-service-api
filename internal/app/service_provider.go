package app

import (
	"context"
	"log"

	"github.com/arpushkarev/note-service-api/internal/config"
	"github.com/arpushkarev/note-service-api/internal/pkg/db"
	"github.com/arpushkarev/note-service-api/internal/repository/note"
	serv "github.com/arpushkarev/note-service-api/internal/service/note"
)

type serviceProvider struct {
	db             db.Client
	configPath     string
	config         config.IConfig
	noteRepository note.Repository
	noteService    *serv.Service
}

func newServiceProvider(configPath string) *serviceProvider {
	return &serviceProvider{
		configPath: configPath,
	}
}

// GetDB gets db
func (s *serviceProvider) GetDB(ctx context.Context) db.Client {
	if s.db == nil {
		cfg, err := s.GetConfig().GetDBConfig()
		if err != nil {
			log.Fatalf("failed to get db config: %s", err.Error())
		}

		dbc, err := db.NewClient(ctx, cfg)
		if err != nil {
			log.Fatalf("can`t connect to db err: %s", err.Error())
		}

		s.db = dbc
	}

	return s.db
}

// GetConfig gets config
func (s *serviceProvider) GetConfig() config.IConfig {
	if s.config == nil {
		cfg, err := config.NewConfig(s.configPath)
		if err != nil {
			log.Fatalf("failed to get config: %s", err.Error())
		}

		s.config = cfg
	}

	return s.config
}

// GetRepository gets repo
func (s *serviceProvider) GetRepository(ctx context.Context) note.Repository {
	if s.noteRepository == nil {
		s.noteRepository = note.NewRepository(s.GetDB(ctx))
	}

	return s.noteRepository
}

// GetNoteService gets service
func (s *serviceProvider) GetNoteService(ctx context.Context) *serv.Service {
	if s.noteService == nil {
		s.noteService = serv.NewService(s.GetRepository(ctx))
	}

	return s.noteService
}
