package message

import (
    "errors"
)

type Application struct {
    ID          string `json:"id"          example:"98q34hiufqh894yh554"`
    Name        string `json:"name"        example:"Data Privacy"`
    Description string `json:"description" example:"Data privacy application description"`
    Type        string `json:"type"        example:"web"`
    Enabled     bool   `json:"enabled"     example:"true"`
}

func (r *Application) Validate() error {
    if r.Name == "" {
        return errors.New("application name is required")
    }

    if r.Type == "" {
        return errors.New("application type is required")
    }

    return nil
}

type OneApplicationRequest struct {
    ID string `json:"id" example:"98q34hiufqh894yh554"`
}

func (r *OneApplicationRequest) Validate() error {
    if r.ID == "" {
        return errors.New("application ID is required")
    }

    return nil
}

type ListApplicationRequest struct {
    Page    uint   `json:"page"    example:"1"`
    Limit   uint   `json:"limit"   example:"10"`
    Type    string `json:"type"    example:"web"`
    Name    string `json:"name"    example:"Data privacy"`
    Enabled string `json:"enabled" example:"1"`
}

func (r *ListApplicationRequest) Validate() error {
    if r.Page == 0 {
        r.Page = 1
    }

    if r.Limit == 0 {
        r.Limit = 10
    }

    return nil
}

type ListApplicationResponse struct {
    Items      []*Application `json:"items"`
    TotalCount uint           `json:"total_count"`
    MaxPage    uint           `json:"max_page"`
    Page       uint           `json:"page"`
    Limit      uint           `json:"limit"`
}
