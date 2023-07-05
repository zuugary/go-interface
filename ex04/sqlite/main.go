package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	ex "github.com/zuugary/go-interface/ex04"
)

func main() {
	db, err := connectToDB()
	if err != nil {
		log.Fatal("Database connection error", err)
	}

	store := ex.NewSqliteStore(db)
	ctr := ex.NewController(store)

	router := gin.Default()
	router.GET("api/v1/:company", ctr.GetCompany)
	router.POST("api/v1/company", ctr.PostCompany)
	router.PUT("api/v1/:company", ctr.UpdateCompany)
	router.DELETE("api/v1/:company", ctr.DeleteCompany)

	log.Fatal(router.Run(":4000"))
}

func connectToDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&ex.Company{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
