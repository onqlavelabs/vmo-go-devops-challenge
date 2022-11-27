package model

import (
    "github.com/golang-jwt/jwt"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseModel struct {
    ID primitive.ObjectID `bson:"_id,omitempty"`
}

type JwtAuthClaim struct {
    jwt.StandardClaims
    Username string `json:"username"`
}
