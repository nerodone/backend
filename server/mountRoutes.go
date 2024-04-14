package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

type Route struct {
	Endpoint     string
	Handler      *chi.Mux
	IsAuthorized bool
}

func (s *Server) MountRoutes(route ...Route) {
	for _, rt := range route {
		if rt.IsAuthorized {
			s.App.Group(func(r chi.Router) {
				r.Use(jwtauth.Verifier(s.JWT.TokenAuth))
				r.Use(jwtauth.Authenticator(s.JWT.TokenAuth))
				r.Mount("/", rt.Handler)
				s.App.Mount(rt.Endpoint, r)
			})
		} else {
			s.App.Mount(rt.Endpoint, rt.Handler)
		}
	}
}
