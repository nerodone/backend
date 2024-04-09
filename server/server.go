package server

import (
	"backend/internal/database"
	"context"
	"database/sql"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

type ApiConfig struct {
	PORT string
	Ctx  context.Context
}
type Server struct {
	ApiConfig *ApiConfig
	JWT       JwtProvider
	App       *chi.Mux
	Db        *database.Queries
}

func New() *Server {
	App := chi.NewRouter()

	App.Use(middleware.Logger)

	db, err := sql.Open("postgres", os.Getenv("XATA_PG"))
	if err != nil {
		panic(err)
	}

	jwt := Init(os.Getenv("JWT_SECRET"))

	return &Server{
		App:       App,
		ApiConfig: &ApiConfig{PORT: os.Getenv("PORT")},
		Db:        database.New(db),
		JWT:       jwt,
	}
}
