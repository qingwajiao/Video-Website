package main

import (
	handles2 "Video-Website/api/handles"
	"Video-Website/api/middleWare"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/scheduler", handles2.CreateUser)

	router.POST("/scheduler/:username", handles2.Login)

	router.GET("/scheduler/:username", handles2.GetUserInfo)

	router.POST("/scheduler/:username/videos", handles2.AddNewVideo)

	router.GET("/scheduler/:username/videos", handles2.ListAllVideos)

	router.DELETE("/scheduler/:username/videos/:vid-id", handles2.DeleteVideo)

	router.POST("/videos/:vid-id/comments", handles2.PostComment)

	router.GET("/videos/:vid-id/comments", handles2.ShowComments)
	return router
}

func main() {
	r := RegisterHandlers()
	m := middleWare.NewMiddleWareHandler(r)

	http.ListenAndServe(":8000", m)
}
