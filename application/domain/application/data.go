package application

import (
    "github.com/dinhtp/vmo-go-devops-challenge/application/database"
    "github.com/dinhtp/vmo-go-devops-challenge/application/message"
    "github.com/dinhtp/vmo-go-devops-challenge/application/model"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func prepareApplicationToResponse(o *model.Application) *message.Application {
    return &message.Application{
        ID:          o.ID.Hex(),
        Name:        o.Name,
        Description: o.Description,
        Type:        o.Type,
        Enabled:     o.Enabled,
    }
}

func prepareApplicationToCreate(o *message.Application) *model.Application {
    return &model.Application{
        BaseModel:   database.BaseModel{ID: primitive.NewObjectID()},
        Name:        o.Name,
        Description: o.Description,
        Enabled:     o.Enabled,
        Type:        o.Type,
    }
}

func prepareApplicationToUpdate(o *message.Application) *model.Application {
    data := prepareApplicationToCreate(o)
    data.ID, _ = primitive.ObjectIDFromHex(o.ID)
    return data
}
