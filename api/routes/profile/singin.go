package profile

import (
	"app/api/config"

	"github.com/goaperture/goaperture/lib/aperture"
)

type SinginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SinginOutput interface {
	string
}

func SinginTest(invoke func(SinginInput)) {
	invoke(SinginInput{})
}

func Singin(input SinginInput, client config.Client) (any, error) {
	if input.Email != "admin" && input.Password != "admin" {
		panic("Не удалось выполнить вход")
	}

	clientId := "555"

	return client.NewJwt(GetPayload(clientId), "refresh_token", clientId, "/profile/refresh", true)
}

func GetPayload(clientId string) config.Payload {
	return config.Payload{
		Id:          clientId,
		Name:        "Admin",
		Email:       "admin@mail.ru",
		Avatar:      "/noimage.svg",
		Permissions: aperture.Permissions{"private"},
	}
}
