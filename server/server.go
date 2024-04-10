package server

import (
	"backend/internal/database"
	"context"
	"database/sql"
	"os"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

type Server struct {
	PORT string
	Ctx  context.Context
	Log  *log.Logger
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

	logger := log.NewWithOptions(os.Stdin, log.Options{
		ReportCaller:    true,
		ReportTimestamp: false,
	})
	return &Server{
		PORT: os.Getenv("PORT"),
		Ctx:  context.Background(),
		Log:  logger,
		App:  App,
		Db:   database.New(db),
		JWT:  jwt,
	}
}
