package model

type Application struct {
    BaseModel   `bson:",inline"`
    Name        string `bson:"name"`
    Description string `bson:"description"`
    Enabled     bool   `bson:"enabled"`
    Type        string `bson:"type"`
}
