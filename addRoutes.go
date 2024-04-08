package main

import "github.com/go-chi/chi/v5"

func (s *server) addRoutes() {
	r := chi.NewRouter()

	// handlers for /api/users
	users := s.userRoutes()
	r.Mount("/users", users)

	s.app.Mount("api/", r)
}
