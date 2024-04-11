package server

import (
	"github.com/go-chi/chi/v5"
)

type Route struct {
	Endpoint string
	Handler  *chi.Mux
}

func (s *Server) MountRoutes(route ...Route) {
	for _, rt := range route {
		s.App.Mount(rt.Endpoint, rt.Handler)
	}
}
