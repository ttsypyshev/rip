package main

import (
	"log"

	"test/internal/api"
)


func main() {
	log.Println("App start")
	api.StartServer()
	log.Println("App down")
}