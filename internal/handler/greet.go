package handler

import (
	"net/http"
)

// all route handlers go here

func (handler *Handler)Greet(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("hello"))
}