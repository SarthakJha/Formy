package handler

import (
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/api/drive/v2"
	"google.golang.org/api/sheets/v4"
)


type Handler struct {
	Db *mongo.Client
	Queue *redis.Client
	SheetsClient *sheets.Service
	DriveClient *drive.Service
}