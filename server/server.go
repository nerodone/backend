package server

import (
	"backend/internal/database"
	"context"
	"database/sql"
	"log/slog"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

type Server struct {
	PORT string
	Ctx  context.Context
	Log  *slog.Logger
	App  *chi.Mux
	Db   *database.Queries
	JWT  JwtProvider
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
		PORT: os.Getenv("PORT"),
		Ctx:  context.Background(),
		Log:  slog.Default(),
		App:  App,
		Db:   database.New(db),
		JWT:  jwt,
	}
}
