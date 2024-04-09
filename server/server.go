package server

import (
	"backend/internal/database"
	"database/sql"
	"os"

	_ "github.com/lib/pq"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
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