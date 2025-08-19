package v1

import "app/api/config"

type PrivateInput struct {
	Token string `json:"token"`
}

type PrivateOutput interface {
	string
}

func PrivateTest(invoke func(PrivateInput)) {
	invoke(PrivateInput{})
}

func Private(input PrivateInput, client config.Client) (any, error) {
	if client.Payload == nil || !client.Payload.Permissions.Has("private") {
		panic("Нет доступа")
	}

	return "hello from private", nil
}
