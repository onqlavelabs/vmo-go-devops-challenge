package database

import (
    "context"
    "github.com/dinhtp/vmo-go-devops-challenge/application/logger"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDatabase struct {
    Uri    string
    Client *mongo.Client
}

func NewMongoDatabase(uri string) *MongoDatabase {
    return &MongoDatabase{
        Uri: uri,
    }
}

func (mg *MongoDatabase) Connect() (*mongo.Client, error) {
    if mg.Client != nil && mg.Client.Ping(context.TODO(), readpref.Primary()) == nil {
        return mg.Client, nil
    }

    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mg.Uri))
    if err != nil {
        return nil, err
    }

    err = client.Ping(context.TODO(), readpref.Primary())
    if err != nil {
        return nil, err
    }

    mg.Client = client
    return client, nil
}

func (mg *MongoDatabase) Disconnect() {
    if mg.Client == nil {
        return
    }

    err := mg.Client.Disconnect(context.TODO())
    if err != nil {
        logger.Log.WithError(err).Error("Disconnect Mongo database connection failed")
    }
}
