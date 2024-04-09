package main

import (
	"backend/server"
	"log"
	"net/http"
)

func main() {
	var s = server.New()
	s.MountRoutes()
	log.Fatal(http.ListenAndServe(":"+s.ApiConfig.PORT, s.App))
}
