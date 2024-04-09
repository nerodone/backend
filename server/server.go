package server

import (
	"backend/internal/database"
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type ApiConfig struct {
	PORT string
	Ctx  context.Context
}
type Server struct {
	ApiConfig *ApiConfig
	App       *chi.Mux
	Db        *database.Queries
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
		App:       App,
		ApiConfig: &ApiConfig{PORT: os.Getenv("PORT")},
		Db:        database.New(db),
	}
}
