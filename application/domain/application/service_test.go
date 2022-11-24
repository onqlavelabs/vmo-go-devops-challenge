package application

import (
    "context"
    "errors"
    "github.com/stretchr/testify/assert"
    "go.mongodb.org/mongo-driver/mongo"
    "net/http"
    "testing"

    "github.com/dinhtp/vmo-go-devops-challenge/application/message"
    "github.com/dinhtp/vmo-go-devops-challenge/application/model"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func TestService_Get(t *testing.T) {
    cases := []struct {
        name   string
        code   int
        method string
        input  *message.OneApplicationRequest
    }{
        {
            name:   "invalid application ID",
            code:   http.StatusBadRequest,
            method: http.MethodGet,
            input:  &message.OneApplicationRequest{ID: "length-24"},
        },
        {
            name:   "application with given ID does not exist",
            code:   http.StatusNotFound,
            method: http.MethodGet,
            input:  &message.OneApplicationRequest{ID: "637ddb4c914039828b39a663"},
        },
        {
            name:   "application found",
            code:   http.StatusOK,
            method: http.MethodGet,
            input:  &message.OneApplicationRequest{ID: "637ddb4c914039828b39a772"},
        },
    }

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            srv := NewService(newMockRepository(tc.method, tc.code))
            result, err := srv.Get(context.Background(), tc.input)

            if tc.code == http.StatusBadRequest {
                assert.NotNil(t, err)
                assert.Error(t, err, primitive.ErrInvalidHex)
            }

            if tc.code == http.StatusNotFound {
                assert.NotNil(t, err)
                assert.Error(t, err, mongo.ErrNoDocuments)
            }

            if tc.code == http.StatusOK {
                assert.NotNil(t, result)
                assert.Equal(t, tc.input.ID, result.ID)
            }
        })
    }
}

func TestService_Create(t *testing.T) {
    cases := []struct {
        name   string
        code   int
        method string
        input  *message.Application
    }{
        {
            name:   "invalid application payload",
            code:   http.StatusBadRequest,
            method: http.MethodPost,
            input:  &message.Application{Description: "mock application"},
        },
        {
            name:   "application created successfully",
            code:   http.StatusCreated,
            method: http.MethodPost,
            input:  &message.Application{Name: "mock-application", Type: "mock", Enabled: true},
        },
    }

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            srv := NewService(newMockRepository(tc.method, tc.code))
            result, err := srv.Create(context.Background(), tc.input)

            if tc.code == http.StatusBadRequest {
                assert.NotNil(t, err)
            }

            if tc.code == http.StatusCreated {
                assert.NotNil(t, result)
                assert.NotEmpty(t, result.ID)
            }
        })
    }
}

type mockRepository struct {
    method string
    code   int
}

func newMockRepository(method string, code int) *mockRepository {
    return &mockRepository{method: method, code: code}
}

func (r *mockRepository) List(ctx context.Context, req *message.ListApplicationRequest) ([]*model.Application, int64, error) {
    return nil, 0, nil
}

func (r *mockRepository) Get(ctx context.Context, id primitive.ObjectID) (*model.Application, error) {
    if r.method != http.MethodGet {
        return nil, nil
    }

    if r.code == http.StatusNotFound {
        return nil, mongo.ErrNoDocuments
    }

    if r.code == http.StatusOK {
        return &model.Application{BaseModel: model.BaseModel{ID: id}, Name: "mock application"}, nil
    }

    return nil, nil
}

func (r *mockRepository) Insert(ctx context.Context, o *model.Application) (*model.Application, error) {
    if r.method != http.MethodPost {
        return nil, nil
    }

    if r.code == http.StatusBadRequest {
        return nil, errors.New("application data is invalid or malformed")
    }

    if r.code == http.StatusCreated {
        return o, nil
    }

    return nil, nil
}

func (r *mockRepository) Update(ctx context.Context, o *model.Application) (*model.Application, error) {
    return nil, nil
}

func (r *mockRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
    return nil
}
