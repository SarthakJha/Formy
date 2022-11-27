package handler

import (
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)


type Handler struct {
	Db *mongo.Client
	Queue *redis.Client
}