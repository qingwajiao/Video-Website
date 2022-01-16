package main

import (
	"Video-Website/scheduler/handlers"
	"Video-Website/scheduler/staskrunner"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	r := httprouter.New()
	r.POST("/video-delete-record/:vid-id", handlers.VidDelRecHandler)
	return r

}

func main() {
	//var c = make(chan string)
	staskrunner.Start()
	r := RegisterHandlers()
	http.ListenAndServe(":9001", r)

	//<- c
}
