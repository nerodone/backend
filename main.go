package main

import "net/http"

func main() {
	s := newSrv()
	s.addRoutes()
	http.ListenAndServe(":"+s.apiCfg.PORT, s.app)
}
