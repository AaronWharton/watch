package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if _, err := io.WriteString(w, "Write string successfully!"); err != nil {
		log.Println(err)
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	if _, err := io.WriteString(w, uname); err != nil {
		log.Println(err)
	}
}
