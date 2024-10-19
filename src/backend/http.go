package backend

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/home", handleHome)
	r.GET("/info/:id", handleInfo)
	r.GET("/app/:id", handleApp)
}

func handleHome(c *gin.Context) {
	query := c.Query("search")
	services := GetLangs()

	var filteredLangs []Lang
	if query != "" {
		filteredLangs = FilterLangsByQuery(services, query)
	} else {
		filteredLangs = services
	}

	c.HTML(http.StatusOK, "services.tmpl", gin.H{
		"Title": "Services",
		"Langs": filteredLangs,
	})
}

func handleInfo(c *gin.Context) {
	id, err := getIDParam(c, "id")
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	info, found := GetLangByID(id)
	if !found {
		c.String(http.StatusNotFound, "Информация о языке с ID %d отсутствует", id)
		return
	}

	c.HTML(http.StatusOK, "information.tmpl", gin.H{
		"Title": info.Name,
		"Info":  info,
		"List":  ParseList(info.List),
	})
}

func handleApp(c *gin.Context) {
	id, err := getIDParam(c, "id")
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	files := GetFilesForProject(id)
	langs := GetLangs()

	c.HTML(http.StatusOK, "applications.tmpl", gin.H{
		"Title": "Project",
		"Files": files,
		"Langs": langs,
	})
}

func getIDParam(c *gin.Context, param string) (int, error) {
	idStr := c.Param(param)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Invalid ID format: %s, error: %v", idStr, err)
		err = errors.New("Не правильный формат id ")
		return 0, err
	}
	return id, nil
}
