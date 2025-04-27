package config

import "github.com/goaperture/goaperture/lib/client"

type Client struct {
	Id          uint               `json:"id"`
	Name        string             `json:"name"`
	Email       string             `json:"email"`
	Avatar      string             `json:"avatar"`
	Permissions client.Permissions `json:"permissions"`
}
