package message

import "errors"

type AuthRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func (r *AuthRequest) Validate() error {
    if r.Username == "" {
        return errors.New("username is required")
    }

    if r.Password == "" {
        return errors.New("password is required")
    }

    return nil
}

type AuthResponse struct {
    BearerToken string `json:"bearer_token"`
    ExpiredIn   int    `json:"expired_in"`
}
