package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Language struct {
	Name        string
	Description string
	ImageURL    string
}

func StartServer() {
	log.Println("Server starting up")

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/image", "./resources")
	r.Static("/static", "./static")

	data := []Language{
		{Name: "Python", Description: "“Объединяет простоту и мощь”", ImageURL: "/static/images/python.png"},
		{Name: "C++", Description: "“Контроль и производительность в одном лице”", ImageURL: "/static/images/cpp.png"},
		{Name: "GO", Description: "“Эффективный для масштабируемых решений”", ImageURL: "/static/images/go.png"},
		{Name: "HTML", Description: "“Основа структуры и содержания веб-страниц”", ImageURL: "/static/images/html.png"},
		{Name: "CSS", Description: "“Создает оформление веб-интерфейсов”", ImageURL: "/static/images/css.png"},
	}

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "services.tmpl", gin.H{
			"title":     "Services",
			"languages": data,
		})
	})

	if err := r.Run(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

	log.Println("Server down")
}
