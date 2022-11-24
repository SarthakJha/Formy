package handler

import "go.mongodb.org/mongo-driver/mongo"


type Handler struct {
	Db *mongo.Client
}