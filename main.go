package main

import (
	"log"
	"net/http"
)

func main() {
	s := newSrv()
	s.addRoutes()
	log.Fatal(http.ListenAndServe(":"+s.apiCfg.PORT, s.app))
}
