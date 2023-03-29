package main

import (
	"atodo/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/user", controllers.CreateUsers)
	router.GET("/users", controllers.GetAll)
	router.GET("/user/:id", controllers.GetByID)
	router.PUT("/user/:id", controllers.UpdateUser)
	router.DELETE("/user/:id", controllers.DeleteUSer)

	err := router.Run(":3000")
	if err != nil {
		return
	}
}
