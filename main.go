package main

import (
	"backend/auth"
	"backend/server"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// only load env file when running locally
	if os.Getenv("KOYEB_APP_NAME") == "" && os.Getenv("RUNTIME") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file", err)
		}
	}

	s := server.New()

	s.MountRoutes(auth.AuthRoutes(s))
	log.Fatal(http.ListenAndServe(":"+s.PORT, s.App))
}
