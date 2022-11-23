package auth

import (
    "context"
    "errors"
    "time"

    "github.com/dinhtp/vmo-go-devops-challenge/application/message"
    "github.com/dinhtp/vmo-go-devops-challenge/application/model"
    "github.com/golang-jwt/jwt"
)

type Service struct {
}

func NewService() *Service {
    return &Service{}
}

func (s *Service) Authenticate(ctx context.Context, r *message.AuthRequest) (*message.AuthResponse, error) {
    if r.Username != DefaultUser || r.Password != DefaultPassword {
        return nil, errors.New("invalid username or password")
    }

    // create jwt Auth claims
    expiredIn := 12 * 60 * 60 // 12 hours
    claims := model.JwtAuthClaim{
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Second * time.Duration(expiredIn)).Unix(),
            IssuedAt:  time.Now().Unix(),
        },
        Username: DefaultUser,
    }

    // create jwt token and sign the token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedToken, err := token.SignedString([]byte(DefaultTokenSecret))
    if err != nil {
        return nil, err
    }

    return &message.AuthResponse{BearerToken: signedToken, ExpiredIn: expiredIn}, nil
}
