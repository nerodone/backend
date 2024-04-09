package main

import (
	"backend/auth"
	"backend/server"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	s := server.New()

	s.MountRoutes(auth.AuthRoutes(s))
	log.Fatal(http.ListenAndServe(":"+s.PORT, s.App))
}
