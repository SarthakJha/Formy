package api

import (
	"github.com/gorilla/mux"
	"github.com/sarthakjha/Formy/internal/handler"
)

func SetupRoutes(router *mux.Router){
	// all api routes go here

	router.HandleFunc("/",handler.Greet)
}