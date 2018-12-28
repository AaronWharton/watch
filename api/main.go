package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middlewareHandler struct {
	r *httprouter.Router
}

func NewMiddlewareHandler(r *httprouter.Router) http.Handler {
	m := middlewareHandler{}
	m.r = r
	return m
}

func (m middlewareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check session
	validateUserSession(r)
	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)

	router.POST("/user/:user_name", Login)

	return router
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddlewareHandler(r)
	http.ListenAndServe("8080", mh)
}
