package model

import (
    "github.com/dinhtp/vmo-go-devops-challenge/application/database"
)

type Application struct {
    database.BaseModel
    Name        string `bson:"name"`
    Description string `bson:"description"`
    Enabled     bool   `bson:"enabled"`
    Type        string `bson:"type"`
}
