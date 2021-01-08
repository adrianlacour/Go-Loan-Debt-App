package main

import (
	"../pkg/controllers"
	"../pkg/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.POST("/loan", controllers.CreateLoan)
	r.GET("/loans", controllers.FindLoans)

	r.Run()
}
