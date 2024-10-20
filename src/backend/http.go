package backend

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app *App) SetupRoutes(r *gin.Engine) {
	r.GET("/home", app.handleHome)
	r.GET("/info/:id", app.handleInfo)
	r.GET("/project/:id", app.handleApp)
	r.POST("/add-service", app.handleAddService)
	r.POST("/upd-project", app.handleUpdateProject)
}

func (app *App) handleHome(c *gin.Context) {
	langID, err := parseQueryParam(c, "langID")
	if err != nil {
		handleError(c, http.StatusBadRequest, errors.New("invalid langID"), err)
		return
	}

	status, err := parseQueryParam(c, "status")
	if err != nil {
		handleError(c, http.StatusBadRequest, errors.New("invalid status"), err)
		return
	}

	query := c.Query("langname")
	filteredLangs, err := app.getFilteredLangs(query)
	if err != nil {
		handleError(c, http.StatusNotFound, errors.New("failed to retrieve language information"), err)
		return
	}

	count, err := app.CountFiles(app.userID)
	if err != nil {
		handleError(c, http.StatusNotFound, errors.New("project was not created"), err)
		return
	}

	projectID, err := findLastDraft(app, app.userID)
	if err != nil {
		handleError(c, http.StatusNotFound, errors.New("project was not created"), err)
		return
	}

	c.HTML(http.StatusOK, "services.tmpl", gin.H{
		"Title":     "Langs",
		"Langs":     filteredLangs,
		"Count":     count,
		"UserID":    app.userID,
		"ProjectID": projectID,
		"LangID":    langID,
		"Status":    status,
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

	project, err := app.GetProjectByID(id)
	if err != nil {
		handleError(c, http.StatusNotFound, errors.New("failed to retrieve language information"), err)
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
		"Project": project,
		"Files": files,
		"Langs": langs,
	})
}

type RequestAdd struct {
	IDUser int `form:"id_user" json:"id_user"`
	IDLang int `form:"id_lang" json:"id_lang"`
}

func (app *App) handleAddService(c *gin.Context) {
	var req RequestAdd

	if err := c.ShouldBind(&req); err != nil {
		handleError(c, http.StatusNotFound, errors.New("invalid data format"), err)
		return
	}

	projectID, err := createDraft(app, req.IDUser)
	if err != nil {
		handleError(c, http.StatusNotFound, errors.New("error creating project"), err)
		return
	}

	if err := app.AddFile(projectID, req.IDLang, req.IDUser); err != nil {
		log.Printf("INFO: Navigating to home page at URL: {URL}/home?langID=%d&status=%d", req.IDLang, 2)
		c.Redirect(http.StatusSeeOther, fmt.Sprintf("/home?langID=%d&status=%d", req.IDLang, 2))
		return
	}

	log.Printf("INFO: Navigating to home page at URL: {URL}/home?langID=%d&status=%d", req.IDLang, 1)
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/home?langID=%d&status=%d", req.IDLang, 1))
}

type RequestUpdate struct {
	IDProject int            `form:"id_project" json:"id_project"`
	Status    int            `form:"status" json:"status"`
	FileCodes map[int]string `form:"file_codes" json:"file_codes"`
}

func (app *App) handleUpdateProject(c *gin.Context) {
	var req RequestUpdate

	if err := c.ShouldBind(&req); err != nil {
		handleError(c, http.StatusNotFound, errors.New("invalid data format"), err)
		return
	}

	if err := app.UpdateProjectStatus(req.IDProject, req.Status); err != nil {
		handleError(c, http.StatusNotFound, errors.New("failed to update project status"), err)
		return
	}

	if err := app.UpdateFilesCode(req.FileCodes); err != nil {
		handleError(c, http.StatusNotFound, errors.New("failed to update file"), err)
		return
	}

	log.Printf("INFO: Navigating to home page at URL: {URL}/home")
	c.Redirect(http.StatusSeeOther, "/home")
}
