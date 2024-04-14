package main

import (
	"backend/auth"
	"backend/server"
	"backend/workspaces"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	s := server.New()
	s.MountRoutes(auth.AuthRoutes(s), workspaces.WorkspacesRoutes(s))
	log.Fatal(http.ListenAndServe(":"+s.PORT, s.App))
}
