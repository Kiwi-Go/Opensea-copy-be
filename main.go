package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rahmanfadhil/gin-bookstore/controllers"
	"github.com/rahmanfadhil/gin-bookstore/models"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/accounts", controllers.FindAccounts)
	r.GET("/accounts/:id", controllers.FindAccount)
	r.POST("/accounts", controllers.CreateAccount)
	r.PATCH("/accounts/:id", controllers.UpdateAccount)
	r.DELETE("/accounts/:id", controllers.DeleteAccount)

	// Run the server
	r.Run()
}
