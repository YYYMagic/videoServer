package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", homeHandler)
	router.GET("/videos/:vid", videoHandler)
	router.GET("/live/:lid", liveHandler)

	return router
}





func main() {
	r := RegisterHandler()
	http.ListenAndServe(":8080", r)
}