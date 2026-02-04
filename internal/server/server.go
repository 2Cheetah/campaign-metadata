package server

import (
	"net/http"

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
	s.Mux.HandleFunc("GET /ping", s.PingHandler)
	s.Mux.HandleFunc("GET /campaigns/{id}/tags", s.GetCampaignTagsHanlder)
}
