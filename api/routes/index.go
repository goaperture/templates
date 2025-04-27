package routes

import "app/api/config"

type IndexInput struct {
}

type IndexOutput interface {
	string
}

func IndexTest(invoke func(IndexInput)) {
	invoke(IndexInput{})
}

func Index(input IndexInput, client *config.Client) (any, error) {
	if !client.Permissions.Has("test", "1234") {
		panic("Нет доступа")
	}

	return "hello world", nil
}
