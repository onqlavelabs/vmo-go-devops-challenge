package application

import (
    "context"
    "github.com/dinhtp/vmo-go-devops-challenge/application/message"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "math"
)

type Service struct {
    repo *Repository
}

func NewService(db *mongo.Client) *Service {
    return &Service{repo: NewRepository(db)}
}

func (s *Service) Get(ctx context.Context, r *message.OneApplicationRequest) (*message.Application, error) {
    primitiveId, err := primitive.ObjectIDFromHex(r.ID)
    if err != nil {
        return nil, err
    }

    result, err := s.repo.Get(ctx, primitiveId)
    if err != nil {
        return nil, err
    }

    return prepareApplicationToResponse(result), nil
}

func (s *Service) List(ctx context.Context, r *message.ListApplicationRequest) (*message.ListApplicationResponse, error) {
    results, totalCount, err := s.repo.List(ctx)
    if err != nil {
        return nil, err
    }

    var list []*message.Application

    for _, result := range results {
        list = append(list, prepareApplicationToResponse(result))
    }

    return &message.ListApplicationResponse{
        Items:      list,
        TotalCount: uint(totalCount),
        MaxPage:    uint(math.Ceil(float64(totalCount) / float64(r.Limit))),
        Page:       r.Page,
        Limit:      r.Limit,
    }, nil
}

func (s *Service) Create(ctx context.Context, r *message.Application) (*message.Application, error) {
    result, err := s.repo.Insert(ctx, prepareApplicationToCreate(r))
    if err != nil {
        return nil, err
    }

    return prepareApplicationToResponse(result), nil
}

func (s *Service) Update(ctx context.Context, r *message.Application) (*message.Application, error) {
    primitiveId, err := primitive.ObjectIDFromHex(r.ID)
    if err != nil {
        return nil, err
    }

    result, err := s.repo.Get(ctx, primitiveId)
    if err != nil {
        return nil, err
    }

    r.ID = result.ID.String()
    updated, err := s.repo.Update(ctx, prepareApplicationToUpdate(r))
    if err != nil {
        return nil, err
    }

    return prepareApplicationToResponse(updated), nil
}

func (s *Service) Delete(ctx context.Context, r *message.OneApplicationRequest) error {
    primitiveId, err := primitive.ObjectIDFromHex(r.ID)
    if err != nil {
        return err
    }

    result, err := s.repo.Get(ctx, primitiveId)
    if err != nil {
        return err
    }

    return s.repo.Delete(ctx, result.ID)
}
