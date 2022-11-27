package application

import (
    "context"
    "math"

    "github.com/dinhtp/vmo-go-devops-challenge/application/logger"
    "github.com/dinhtp/vmo-go-devops-challenge/application/message"
    "github.com/sirupsen/logrus"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
    repo RepositoryInterface
}

func NewService(repo RepositoryInterface) *Service {
    return &Service{repo: repo}
}

func (s *Service) Get(ctx context.Context, r *message.OneApplicationRequest) (*message.Application, error) {
    primitiveId, err := primitive.ObjectIDFromHex(r.ID)
    if err != nil {
        return nil, err
    }

    result, err := s.repo.Get(ctx, primitiveId)
    if err != nil {
        logger.Log.WithFields(logrus.Fields{
            "package": "application",
            "method":  "Get",
            "request": r,
        }).WithError(err).Error("get application failed")
        return nil, err
    }

    return prepareApplicationToResponse(result), nil
}

func (s *Service) List(ctx context.Context, r *message.ListApplicationRequest) (*message.ListApplicationResponse, error) {
    results, totalCount, err := s.repo.List(ctx, r)
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
    createData := prepareApplicationToCreate(r)
    result, err := s.repo.Insert(ctx, createData)
    if err != nil {
        logger.Log.WithFields(logrus.Fields{
            "package": "application",
            "method":  "Create",
            "request": r,
            "data":    createData,
        }).WithError(err).Error("create new application failed")
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
        logger.Log.WithFields(logrus.Fields{
            "package": "application",
            "method":  "Update",
            "request": r,
        }).WithError(err).Error("get existing application failed")
        return nil, err
    }

    r.ID = result.ID.Hex()
    updateData := prepareApplicationToUpdate(r)
    updated, err := s.repo.Update(ctx, prepareApplicationToUpdate(r))
    if err != nil {
        logger.Log.WithFields(logrus.Fields{
            "package": "application",
            "method":  "Update",
            "request": r,
            "data":    updateData,
        }).WithError(err).Error("update new application failed")
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
        logger.Log.WithFields(logrus.Fields{
            "package": "application",
            "method":  "Delete",
            "request": r,
        }).WithError(err).Error("delete application failed")
        return err
    }

    return s.repo.Delete(ctx, result.ID)
}
