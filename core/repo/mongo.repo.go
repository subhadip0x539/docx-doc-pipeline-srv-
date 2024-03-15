package repo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type IMongoRepo interface {
}

type MongoRepo struct {
	client   *mongo.Client
	database string
}

func NewMongoRepo(client *mongo.Client, database string) IMongoRepo {
	return &MongoRepo{client: client, database: database}
}
