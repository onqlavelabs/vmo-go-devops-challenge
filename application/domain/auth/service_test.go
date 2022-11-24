package auth

import (
    "context"
    "testing"

    "github.com/dinhtp/vmo-go-devops-challenge/application/message"
    "github.com/stretchr/testify/assert"
)

func TestService_Authenticate(t *testing.T) {
    cases := []struct {
        name     string
        errorMsg string
        input    *message.AuthRequest
    }{
        {
            name:     "authenticate failed",
            input:    &message.AuthRequest{Username: "root", Password: "not-root"},
            errorMsg: "invalid username or password",
        },
        {
            name:  "authenticate successfully",
            input: &message.AuthRequest{Username: "root", Password: "root"},
        },
    }

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            srv := NewService()
            output, err := srv.Authenticate(context.Background(), tc.input)

            if tc.errorMsg != "" {
                assert.Error(t, err, tc.errorMsg)
            } else {
                assert.NotEmpty(t, output)
                assert.NotEmpty(t, output.BearerToken)
                assert.Equal(t, ExpiredTime, output.ExpiredIn)
            }
        })
    }
}
