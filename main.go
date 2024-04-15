package main

import (
	"backend/auth"
	"backend/server"
	"log"
	"net/http"
	"os"

	_ "backend/docs"

	"github.com/joho/godotenv"
)

// @title			NeroDone api
// @version		1.0
// @description	This is a sample server celler server.
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
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
