package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Company struct {
	Name    string `json:"name"`
	Created string `json:"created"`
	Product string `json:"product"`
}

var (
	db  *gorm.DB
	err error
)

func DBConnection() error {
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&Company{})
	if err != nil {
		return err
	}

	return nil
}

func GetCompany(ctx *gin.Context) {
	var company Company
	name := ctx.Param("company")
	if err := db.Where("name= ?", name).First(&company).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"Failed": "Company not found"})
		return
	}
	ctx.JSON(http.StatusOK, company)
}

func PostCompany(ctx *gin.Context) {
	var company Company
	if err := ctx.ShouldBindJSON(&company); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error Decoding JSON"})
		return
	}
	newCompany := Company{
		Name:    company.Name,
		Created: company.Created,
		Product: company.Product,
	}

	if err := db.Create(&newCompany).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database Migration Error"})
	}
}

func UpdateCompany(ctx *gin.Context) {
	var company Company

	name := ctx.Param("company")

	if err = db.Where("name = ?", name).First(&company).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	err = ctx.BindJSON(&company)
	if err != nil {
		return
	}
	db.Save(&company)

	ctx.JSON(http.StatusOK, company)
}

func DeleteCompany(ctx *gin.Context) {
	var company Company
	name := ctx.Param("company")
	if err := db.Where("name = ?", name).Delete(&company).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Company Deleted"})
}

func main() {
	err := DBConnection()
	if err != nil {
		log.Fatal("Database connection error", err)
	}
	router := gin.Default()
	router.GET("api/v1/:company", GetCompany)
	router.POST("api/v1/company", PostCompany)
	router.PUT("api/v1/:company", UpdateCompany)
	router.DELETE("api/v1/:company", DeleteCompany)

	log.Fatal(router.Run(":4000"))
}
