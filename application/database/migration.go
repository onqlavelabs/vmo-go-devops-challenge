package database

import (
    "context"

    "github.com/dinhtp/vmo-go-devops-challenge/application/message"
    "github.com/dinhtp/vmo-go-devops-challenge/application/model"
    "github.com/dinhtp/vmo-go-devops-challenge/application/util"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func Migrate(c *mongo.Client) error {
    filter := primitive.M{"name": CollectionApplications}
    result, err := c.Database(Name).ListCollectionNames(context.Background(), filter)
    if err != nil {
        return err
    }

    if len(result) > 0 {
        return nil
    }

    opts := options.CreateCollection().SetValidator(bson.M{"$jsonSchema": ApplicationSchema})
    err = c.Database(Name).CreateCollection(context.Background(), CollectionApplications, opts)
    if err != nil {
        return err
    }

    return nil
}

func Seed(c *mongo.Client) error {
    var seedItems []interface{}
    var data []*message.Application

    err := util.ReadJsonFile("/app/db.json", &data)
    if err != nil {
        return err
    }

    if data == nil || len(data) == 0 {
        return nil
    }

    for _, datum := range data {
        seedItems = append(seedItems, &model.Application{
            BaseModel:   model.BaseModel{ID: primitive.NewObjectID()},
            Name:        datum.Name,
            Description: datum.Description,
            Enabled:     datum.Enabled,
            Type:        datum.Type,
        })
    }

    _, err = c.Database(Name).Collection(CollectionApplications).InsertMany(context.Background(), seedItems)
    if err != nil {
        return err
    }

    return nil
}
