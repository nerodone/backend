package auth

import (
	"backend/server"

	"github.com/go-chi/chi/v5"
)

func authRouter(s *server.Server) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/signup", signup(s))
	r.Post("/login", login(s))
	r.Post("/logout", logout(s))

	return r
}

func AuthRoutes(s *server.Server) server.Route {
	return server.Route{
		Endpoint: "/auth",
		Handler:  authRouter(s),
	}
}
