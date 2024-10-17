package api

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	log.Println("Server starting up")

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/image", "./resources")
	r.Static("/static", "./static")

	r.GET("/home", func(c *gin.Context) {
		query := c.Query("search")

		services := GetServices()
		var filteredLangs []Lang

		if query == "" {
			filteredLangs = services
		} else {
			for _, lang := range services {
				if strings.Contains(strings.ToLower(lang.Name), strings.ToLower(query)) {
					filteredLangs = append(filteredLangs, lang)
				}
			}
		}

		c.HTML(http.StatusOK, "services.tmpl", gin.H{
			"Title": "Services",
			"Langs": filteredLangs,
		})
	})

	r.GET("/info/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		services := GetServices()
		if err != nil || id < 0 || id >= len(services) {
			c.String(http.StatusNotFound, "Страница не найдена")
			return
		}
		info := services[id]

		c.HTML(http.StatusOK, "information.tmpl", gin.H{
			"Title": info.Name,
			"Info":  info,
			"List":  parseList(info.List),
		})
	})

	r.GET("/app/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		application := GetApplications()
		if err != nil || id < 0 || id >= len(application) {
			c.String(http.StatusNotFound, "Страница не найдена")
			return
		}
		app := GetFilesForProject(id)
		langs := GetLangsForProject(app, id)

		c.HTML(http.StatusOK, "applications.tmpl", gin.H{
			"Title": "Applications",
			"App":   app,
			"Lang":  langs[0],
		})
	})

	// r.GET("/app", func(c *gin.Context) {
	// 	id := FindMaxProjectID()
	// 	app := GetFilesForProject(id)
	// 	langs := GetLangsForProject(app, id)

	// 	c.HTML(http.StatusOK, "applications.tmpl", gin.H{
	// 		"Title": "Applications",
	// 		"App":   app,
	// 		"Lang":  langs[0],
	// 	})
	// })

	if err := r.Run(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

	log.Println("Server down")
}
