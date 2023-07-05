package ex04

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	store Store
}

func NewController(s Store) *Controller {
	return &Controller{s}
}

func (c *Controller) GetCompany(ctx *gin.Context) {
	name := ctx.Param("company")

	company, err := c.store.Get(name)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"Failed": "Company not found"})
		return
	}

	ctx.JSON(http.StatusOK, company)
}

func (c *Controller) PostCompany(ctx *gin.Context) {
	var company Company
	if err := ctx.ShouldBindJSON(&company); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error Decoding JSON"})
		return
	}
	newCompany := &Company{
		Name:    company.Name,
		Created: company.Created,
		Product: company.Product,
	}

	if err := c.store.Insert(newCompany); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database Migration Error"})
	}
}

func (c *Controller) UpdateCompany(ctx *gin.Context) {
	name := ctx.Param("company")

	company, err := c.store.Get(name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	_, err = c.store.Update(company)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Update error!"})
		return
	}

	err = ctx.BindJSON(&company)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, company)
}

func (c *Controller) DeleteCompany(ctx *gin.Context) {
	name := ctx.Param("company")

	if err := c.store.Delete(name); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Company Deleted"})
}
