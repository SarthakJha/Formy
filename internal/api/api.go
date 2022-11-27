package api

import (
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/sarthakjha/Formy/internal/handler"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/api/drive/v2"
	"google.golang.org/api/sheets/v4"
)

func SetupRoutes(router *mux.Router, dbClient *mongo.Client, queueClient *redis.Client, sheetsClient *sheets.Service, driveClient *drive.Service){
	handler := handler.Handler{
		Db: dbClient,
		Queue: queueClient,
		SheetsClient: sheetsClient,
		DriveClient: driveClient,
	}

	subRouter := router.PathPrefix("/api/"+os.Getenv("API_VERSION")).Subrouter()

	subRouter.HandleFunc("/create-form",handler.CreateForm)
	subRouter.HandleFunc("/submit-response",handler.CreateResponse)
	
}