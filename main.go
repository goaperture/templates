package main

import (
	"app/api"
	"app/api/config"
	"log"
)

func main() {
	client := config.Client{
		Id:   1,
		Name: "123",
	}

	log.Println(client)

	log.Println("Server run on http://localhost:8080")
	api.Serve(8080, nil)
}
