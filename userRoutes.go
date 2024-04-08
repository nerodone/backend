package main

import "github.com/go-chi/chi/v5"

func (s *server) userRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/signup", handlerUserSignUp)

	return r
}
