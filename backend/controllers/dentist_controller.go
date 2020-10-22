package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/B5916177/app/ent"
	"github.com/B5916177/app/ent/dentist"
	"github.com/gin-gonic/gin"
)

type DentistController struct {
	client *ent.Client
	router gin.IRouter
}

type Dentist struct {
	DentistName   string
	DentistEmail  string
	DentistPhone  int
}

// CreateDentist handles POST requests for adding dentist entities
// @Summary Create dentist
// @Description Create dentist
// @ID create-dentist
// @Accept   json
// @Produce  json
// @Param dentist body ent.Dentist true "Dentist entity"
// @Success 200 {object} ent.Dentist
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dentists [post]
func (ctl *DentistController) CreateDentist(c *gin.Context) {
	obj := Dentist{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "dentist binding failed",
		})
		return
	}

	d, err := ctl.client.Dentist.
		Create().
		SetDentistName(obj.DentistName).
		SetDentistEmail(obj.DentistEmail).
		SetDentistPhone(obj.DentistPhone).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, d)
}

// GetDentist handles GET requests to retrieve a dentist entity
// @Summary Get a dentist entity by ID
// @Description get dentist by ID
// @ID get-dentist
// @Produce  json
// @Param id path int true "Dentist ID"
// @Success 200 {object} ent.Dentist
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dentists/{id} [get]
func (ctl *DentistController) GetDentist(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	d, err := ctl.client.Dentist.
		Query().
		Where(dentist.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, d)
}

// ListDentist handles request to get a list of dentist entities
// @Summary List dentist entities
// @Description list dentist entities
// @ID list-dentist
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Dentist
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dentists [get]
func (ctl *DentistController) ListDentist(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	dentists, err := ctl.client.Dentist.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, dentists)
}

// DeleteDentist handles DELETE requests to delete a dentist entity
// @Summary Delete a dentist entity by ID
// @Description get dentist by ID
// @ID delete-dentist
// @Produce  json
// @Param id path int true "Dentist ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dentist/{id} [delete]
func (ctl *DentistController) DeleteDentist(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Dentist.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// NewDentistController creates and registers handles for the dentist controller
func NewDentistController(router gin.IRouter, client *ent.Client) *DentistController {
	dc := &DentistController{
		client: client,
		router: router,
	}

	dc.register()

	return dc

}

func (ctl *DentistController) register() {
	dentists := ctl.router.Group("/dentists")

	dentists.POST("", ctl.CreateDentist)
	dentists.GET(":id", ctl.GetDentist)
	dentists.GET("", ctl.ListDentist)
	dentists.DELETE("", ctl.DeleteDentist)

}