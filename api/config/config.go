package config

import (
	"github.com/goaperture/goaperture/lib/aperture"
)

type Payload struct {
	Id          string               `json:"id"`
	Name        string               `json:"name"`
	Email       string               `json:"email"`
	Avatar      string               `json:"avatar"`
	Permissions aperture.Permissions `json:"permissions"`
}

type Client = aperture.Client[Payload]
