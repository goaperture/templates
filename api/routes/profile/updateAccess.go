package profile

import (
	"app/api/config"
)

type UpdateAccessInput struct {
	Token string `json:"token"`
}

type UpdateAccessOutput interface {
	string
}

func UpdateAccessTest(invoke func(UpdateAccessInput)) {
	invoke(UpdateAccessInput{})
}

func UpdateAccess(input UpdateAccessInput, client config.Client) (any, error) {
	// clientId, err := jwt.DecodeRefreshToken(input.Token)
	// if err != nil {
	// 	panic("Не удалось выпонить вход")
	// }

	// newClient := config.Client{
	// 	Id:          clientId,
	// 	Name:        "admin",
	// 	Email:       "admin",
	// 	Avatar:      "no-image.svg",
	// 	Permissions: ac.Permissions{"all"},
	// }

	// return jwt.NewAccessToken(newClient)
}
