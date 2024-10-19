package backend

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app *App) SetupRoutes(r *gin.Engine) {
	r.GET("/home", app.handleHome)
	r.GET("/info/:id", app.handleInfo)
	r.GET("/project/:id", app.handleApp)
}

func (app *App) handleHome(c *gin.Context) {
	query := c.Query("langname")
	services, err := app.GetLangs()
	if err != nil {
		handleError(c, http.StatusNotFound, errors.New("failed to retrieve language information"), err)
		return
	}

	var filteredLangs []DbLang
	if query != "" {
		filteredLangs, err = app.FilterLangsByQuery(query)
	} else {
		filteredLangs = services
	}
	if err != nil {
		handleError(c, http.StatusNotFound, errors.New("failed to filter languages by query"), err)
		return
	}

	c.HTML(http.StatusOK, "services.tmpl", gin.H{
		"Title": "Langs",
		"Langs": filteredLangs,
	})
}

func (app *App) handleInfo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		handleError(c, http.StatusNotFound, errors.New("invalid language ID"), err)
		return
	}

	info, err := app.GetLangByID(id)
	if err != nil {
		handleError(c, http.StatusNotFound, errors.New("language information not available"), err)
		return
	}

	c.HTML(http.StatusOK, "information.tmpl", gin.H{
		"Title": info.Name,
		"Info":  info,
		"List":  ParseList(info.List),
	})
}

func (app *App) handleApp(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		handleError(c, http.StatusNotFound, errors.New("invalid project ID"), err)
		return
	}

	files, err := app.GetFilesForProject(id)
	if err != nil {
		handleError(c, http.StatusNotFound, errors.New("failed to retrieve project files"), err)
		return
	}
	langs, err := app.GetLangs()
	if err != nil {
		handleError(c, http.StatusNotFound, errors.New("failed to retrieve language information"), err)
		return
	}

	c.HTML(http.StatusOK, "applications.tmpl", gin.H{
		"Title": "Project",
		"Files": files,
		"Langs": langs,
	})
}
