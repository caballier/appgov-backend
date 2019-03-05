package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ShowApplication godoc
// @Summary Show an application
// @Description get application by ID
// @Tags applications
// @Accept  json
// @Produce  json
// @Param id path int true "Application ID"
// @Success 200 {object} model.Application
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /{id} [get]
func (c *Controller) ShowApplication(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, id)
}

// ListApplications godoc
// @Summary List applications
// @Description get applications
// @Tags applications
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Application
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router / [get]
func (c *Controller) ListApplications(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "OK")
}

// AddApplication godoc
// @Summary Add an application
// @Description add by json application
// @Tags applications
// @Accept  json
// @Produce  json
// @Param account body model.AddApplication true "Add application"
// @Success 200 {object} model.Application
// @Failure 400 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router / [post]
func (c *Controller) AddApplication(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "OK")
}

// UpdateApplication godoc
// @Summary Update an application
// @Description Update by json application
// @Tags applications
// @Accept  json
// @Produce  json
// @Param  id path int true "Application ID"
// @Param  application body model.UpdateApplication true "Update application"
// @Success 200 {object} model.Application
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /{id} [put]
func (c *Controller) UpdateApplication(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, id)
}

// DeleteApplication godoc
// @Summary Delete an application
// @Description Delete by application ID
// @Tags applications
// @Accept  json
// @Produce  json
// @Param  id path int true "Application ID"
// @Success 204 {object} model.Application
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /{id} [delete]
func (c *Controller) DeleteApplication(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, id)
}
