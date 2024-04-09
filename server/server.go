package server

import (
	"backend/internal/database"
	"database/sql"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type ApiConfig struct {
	PORT string
}
type Server struct {
	ApiConfig *ApiConfig
	App       *chi.Mux
	Db        *database.Queries
}

func New() *Server {
	godotenv.Load()
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
