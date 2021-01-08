package main

import (
	"../pkg/controllers"
	"../pkg/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/loans", controllers.FindLoans)
	r.GET("/loans/:id", controllers.FindLoan)
	r.POST("/loans", controllers.CreateLoan)

	r.Run()
}
