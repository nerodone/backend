package server

import "github.com/go-chi/chi/v5"

type Route struct {
	Endpoint string
	Handler  *chi.Mux
}

func (s *Server) MountRoutes(route ...Route) {
	r := chi.NewRouter()

	for _, rt := range route {
		r.Mount(rt.Endpoint, rt.Handler)
	}

	s.App.Mount("/", r)
}
