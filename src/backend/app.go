package backend

import (
	"log"

	"github.com/gin-gonic/gin"
)

func StartServer() error {
	log.Println("Server starting up")

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	SetupRoutes(r)

	if err := r.Run(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
		return err
	}

	log.Println("Server stopped")
	return nil
}
