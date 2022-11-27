package api

import (
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/sarthakjha/Formy/internal/handler"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/api/sheets/v4"
)

func SetupRoutes(router *mux.Router, dbClient *mongo.Client, queueClient *redis.Client, sheetsClient *sheets.Service){
	// all api routes go here
	 
	handler := handler.Handler{
		Db: dbClient,
		Queue: queueClient,
		SheetsClient: sheetsClient,
	}
	// create handler obj here then pass it through the routes
	// fill handler struct here, with values from main through args

	subRouter := router.PathPrefix("/api/"+os.Getenv("API_VERSION")).Subrouter()


	subRouter.HandleFunc("/create-form",handler.CreateForm)
	subRouter.HandleFunc("/submit-response",handler.CreateResponse)
	
	/*
		1. create form route (admin) (mininal traffic)
			1.1 take in questions
			1.2 take required feature - sheets or SMS 

		2. submit response route (public) (expected max traffic)
			2.1 store response in data store
				2.1.1 worker pool or general
			2.2 forward data to opted service(sms/sheet) to carry their processing
				2.2.1 syncronous/async
	*/
}