package application

import (
    "context"
    "github.com/dinhtp/vmo-go-devops-challenge/application/database"

    "github.com/dinhtp/vmo-go-devops-challenge/application/message"
    "github.com/dinhtp/vmo-go-devops-challenge/application/model"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
    db *mongo.Client
}

func NewRepository(db *mongo.Client) *Repository {
    return &Repository{db: db}
}

func (r *Repository) List(ctx context.Context, req *message.ListApplicationRequest) ([]*model.Application, int64, error) {
    var results []*model.Application

    query := r.buildQuery()
    filter := r.buildListFilter(req)
    opts := options.Find().SetLimit(int64(req.Limit)).SetSkip(int64(req.Limit) * (int64(req.Page) - 1))

    totalCount, err := query.CountDocuments(ctx, filter)
    if err != nil {
        return nil, 0, err
    }

    cursor, err := query.Find(ctx, filter, opts)
    if err != nil {
        return nil, 0, err
    }

    if err := cursor.All(ctx, &results); err != nil {
        return nil, 0, err
    }

    return results, totalCount, nil
}

func (r *Repository) Get(ctx context.Context, id primitive.ObjectID) (*model.Application, error) {
    var result *model.Application

    err := r.buildQuery().FindOne(ctx, primitive.M{"_id": id}).Decode(&result)
    if err != nil {
        return nil, err
    }

    return result, nil
}

func (r *Repository) Insert(ctx context.Context, o *model.Application) (*model.Application, error) {
    _, err := r.buildQuery().InsertOne(ctx, o)
    if err != nil {
        return nil, err
    }

    return o, nil
}

func (r *Repository) Update(ctx context.Context, o *model.Application) (*model.Application, error) {
    _, err := r.buildQuery().UpdateOne(ctx, primitive.M{"_id": o.ID}, primitive.M{"$set": o})
    if err != nil {
        return nil, err
    }

    return o, nil
}

func (r *Repository) Delete(ctx context.Context, id primitive.ObjectID) error {
    _, err := r.buildQuery().DeleteOne(ctx, primitive.M{"_id": id})
    return err
}

func (r *Repository) buildQuery() *mongo.Collection {
    return r.db.Database(database.Name).Collection(database.CollectionApplications)
}

func (r *Repository) buildListFilter(req *message.ListApplicationRequest) primitive.M {
    filter := primitive.M{}
    if req.Name != "" {
        filter["name"] = req.Name
    }

    if req.Type != "" {
        filter["type"] = req.Type
    }

    if req.Enabled == "1" {
        filter["enabled"] = true
    }

    if req.Enabled == "0" {
        filter["enabled"] = false
    }

    return filter
}
