package migration

import (
    "context"

    "github.com/dinhtp/vmo-go-devops-challenge/application/database"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func Migrate(c *mongo.Client) error {
    opts := options.CreateCollection().SetValidator(bson.M{"$jsonSchema": database.ApplicationSchema})

    err := c.Database(database.Name).CreateCollection(context.Background(), database.CollectionApplications, opts)
    if err != nil {
        return err
    }

    return nil
}
