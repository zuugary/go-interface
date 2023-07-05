package main

import (
	"log"

	"github.com/gin-gonic/gin"
	ex "github.com/zuugary/go-interface/ex04"
)

func main() {
	store := ex.NewInmemStore()
	ctr := ex.NewController(store)

	router := gin.Default()
	router.GET("api/v1/:company", ctr.GetCompany)
	router.POST("api/v1/company", ctr.PostCompany)
	router.PUT("api/v1/:company", ctr.UpdateCompany)
	router.DELETE("api/v1/:company", ctr.DeleteCompany)

	log.Fatal(router.Run(":4000"))
}
