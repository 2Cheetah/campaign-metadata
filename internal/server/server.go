package server

import (
	"net/http"
	"os"

	"github.com/2Cheetah/campaign-metadata/internal/service"
)

type Server struct {
	Service *service.Service
	Mux     *http.ServeMux
	Addr    string
}

func NewServer(service *service.Service) *Server {
	s := &Server{
		Service: service,
		Mux:     http.NewServeMux(),
		Addr:    ":8080",
	}
	s.RegisterHandlers()
	return s
}

func (s *Server) RegisterHandlers() {
	// Serve static files
	fsys := os.DirFS("web/static")
	fs := http.FileServerFS(fsys)
	s.Mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	s.Mux.HandleFunc("GET /ping", s.PingHandler)
	s.Mux.HandleFunc("GET /campaigns/{id}/tags", s.GetCampaignTagsHanlder)
	s.Mux.HandleFunc("GET /index", s.IndexHandler)
}
