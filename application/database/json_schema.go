package database

import (
    "go.mongodb.org/mongo-driver/bson"
)

var ApplicationSchema = bson.M{
    "bsonType":             "object",
    "required":             []string{"name", "enabled", "type"},
    "additionalProperties": false,
    "properties": bson.M{
        "_id": bson.M{
            "bsonType": "objectId",
        },
        "name": bson.M{
            "bsonType":    "string",
            "description": "required string",
            "minLength":   1,
            "maxLength":   255,
        },
        "description": bson.M{
            "bsonType":    "string",
            "description": "optional description",
            "minLength":   0,
            "maxLength":   255,
        },
        "enabled": bson.M{
            "bsonType":    "bool",
            "description": "required boolean",
        },
        "type": bson.M{
            "bsonType":    "string",
            "description": "required string",
            "minLength":   1,
            "maxLength":   50,
        },
    },
}
