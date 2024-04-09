package auth

import "github.com/go-chi/chi/v5"

func AuthRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/signup", handlerUserSignUp)

	return r
}
