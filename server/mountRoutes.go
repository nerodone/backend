package server

import "github.com/go-chi/chi/v5"
import "backend/auth"

func (s *Server) MountRoutes() {
	r := chi.NewRouter()

	r.Mount("/auth", auth.AuthRouter())

	s.App.Mount("/", r)
}
