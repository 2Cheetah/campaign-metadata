package app

import (
	"context"
	"log"
	"net/http"

	"github.com/2Cheetah/campaign-metadata/internal/repository"
	"github.com/2Cheetah/campaign-metadata/internal/server"
	"github.com/2Cheetah/campaign-metadata/internal/service"
)

type App struct {
	Server *server.Server
}

func NewApp(ctx context.Context) *App {
	supabaseDB, err := repository.NewSupabaseDB(ctx)
	if err != nil {
		log.Fatalf("couldn't instantiate Supabase DB, error: %v", err)
	}
	service := service.NewService(supabaseDB)
	server := server.NewServer(service)
	return &App{
		Server: server,
	}
}

func (a *App) MustRun() {
	if err := http.ListenAndServe(a.Server.Addr, a.Server.Mux); err != nil {
		log.Fatalf("couldn't start http server, error: %v", err)
	}
}
