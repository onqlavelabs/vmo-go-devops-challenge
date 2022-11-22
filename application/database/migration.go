package database

import (
    "context"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func Migrate(c *mongo.Client) error {
    opts := options.CreateCollection().SetValidator(bson.M{"$jsonSchema": ApplicationSchema})

    err := c.Database(Name).CreateCollection(context.Background(), CollectionApplications, opts)
    if err != nil {
        return err
    }

    return nil
}
