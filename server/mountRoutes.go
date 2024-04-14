package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

type Route struct {
	Endpoint     string
	Handler      *chi.Mux
	isAuthorized bool
}

func (s *Server) MountRoutes(route ...Route) {
	for _, rt := range route {
		if rt.isAuthorized {
			rt.Handler.Use(jwtauth.Verifier(s.JWT.TokenAuth))
			rt.Handler.Use(jwtauth.Authenticator(s.JWT.TokenAuth))
			s.App.Mount(rt.Endpoint, rt.Handler)
		}
	}
}
