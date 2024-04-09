package server

import (
	"backend/internal/database"
	"context"
	"database/sql"
	"log"
	"log/slog"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Server struct {
	PORT       string
	Ctx        context.Context
	Log        *slog.Logger
	App        *chi.Mux
	Db         *database.Queries
	JWT_SECRET []byte
}

func New() *Server {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	App := chi.NewRouter()

	App.Use(middleware.Logger)

	db, err := sql.Open("postgres", os.Getenv("XATA_PG"))
	if err != nil {
		panic(err)
	}

	return &Server{
		App:  App,
		PORT: os.Getenv("PORT"),
		Ctx:  context.Background(),
		Db:   database.New(db),
		Log:  slog.Default(),
	}
}
