package profile

import (
	"app/api/config"

	"github.com/goaperture/goaperture/lib/aperture"
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
	return client.RefreshJwt("refresh_token", func(refreshId string) config.Payload {
		return config.Payload{
			Id:          refreshId,
			Name:        "Admin",
			Email:       "admin@mail.ru",
			Avatar:      "/noimage.svg",
			Permissions: aperture.Permissions{"noo"},
		}
	})
}
