package profile

import (
	"app/api/config"
)

type SingoutInput struct {
}

type SingoutOutput interface {
	string
}

func SingoutTest(invoke func(SingoutInput)) {
	invoke(SingoutInput{})
}

func Singout(input SingoutInput, client config.Client) (any, error) {
	client.RemoveJwt("refresh_token", "/profile/refresh", true)

	return "OK", nil
}
