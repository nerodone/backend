package main

import (
	"database/sql"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	PORT string
}
type server struct {
	apiCfg *apiConfig
	app    *chi.Mux
	db     *database.Queries
}

func newSrv() *server {
	godotenv.Load()
	app := chi.NewRouter()

	db, err := sql.Open("postgres", os.Getenv("XATA_PG"))
	if err != nil {
		panic(err)
	}

	return &server{
		app:    app,
		apiCfg: &apiConfig{PORT: os.Getenv("APIPORT")},
		db:     database.New(db),
	}
}
