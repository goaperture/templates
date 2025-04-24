package main

import (
	"app/api"
	"log"
)

func main() {
	log.Println("Server run on http:/localhost:8080")
	api.Serve(8080, nil)
}
