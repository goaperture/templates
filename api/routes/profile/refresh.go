package profile

import (
	"app/api/config"
)

type RefreshInput struct {
}

type RefreshOutput interface {
	string
}

func RefreshTest(invoke func(RefreshInput)) {
	invoke(RefreshInput{})
}

func Refresh(input RefreshInput, client config.Client) (any, error) {
	return client.RefreshJwt("refresh_token", GetPayload)
}
