package respository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDataBase(url string) (*mongo.Client, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client,err := mongo.Connect(ctx,options.Client().ApplyURI(url))

	return client,err
}