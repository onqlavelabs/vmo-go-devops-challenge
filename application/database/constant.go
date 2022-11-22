package database

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

const (
    Name = "programming_challenge"

    CollectionApplications = "applications"
)

type BaseModel struct {
    ID primitive.ObjectID `bson:"_id,omitempty"`
    //CreatedAt *JSONTime          `json:"created_at" bson:"created_at,omitempty"`
    //UpdatedAt *JSONTime          `json:"updated_at" bson:"updated_at,omitempty"`
    //DeletedAt *JSONTime          `json:"deleted_at" bson:"deleted_at,omitempty"`
}
