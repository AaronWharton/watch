package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)

	router.POST("/user/:user_name", Login)
	
	return router
}

func main() {
	r := RegisterHandlers()
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Println(err)
	}
}
