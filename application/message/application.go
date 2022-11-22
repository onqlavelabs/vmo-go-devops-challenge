package message

type Application struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    Type        string `json:"type"`
    Enabled     bool   `json:"enabled"`
}

type OneApplicationRequest struct {
    ID string `json:"id"`
}

type ListApplicationRequest struct {
    Page    uint   `json:"page"`
    Limit   uint   `json:"limit"`
    Type    string `json:"type"`
    Name    string `json:"name"`
    Enabled string `json:"enabled"`
}

type ListApplicationResponse struct {
    Items      []*Application `json:"items"`
    TotalCount uint           `json:"total_count"`
    MaxPage    uint           `json:"max_page"`
    Page       uint           `json:"page"`
    Limit      uint           `json:"limit"`
}
