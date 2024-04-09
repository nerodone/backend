package main

import (
	"backend/auth"
	"backend/server"
	"log"
	"net/http"
)

func main() {
	var s = server.New()
	s.MountRoutes(auth.AuthRoutes(s))
	log.Fatal(http.ListenAndServe(":"+s.ApiConfig.PORT, s.App))
}
