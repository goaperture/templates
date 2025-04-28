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
	if client.Payload == nil {

	}

	return client.Payload.Avatar, nil
}
