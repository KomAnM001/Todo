package main

import (
	"atodo/controllers"
	"atodo/storage/todo"
	"github.com/gin-gonic/gin"
)

func main() {

	strg, err := todo.NewStore()
	if err != nil {
		panic("error while db conn" + err.Error())
	}
	tdSrvcs := controllers.NewServiceTodo(strg)

	router := gin.Default()
	router.POST("/user", tdSrvcs.CreateUsers)
	router.GET("/user/:id", tdSrvcs.GetByID)
	router.GET("/users", tdSrvcs.GetAll)
	router.PUT("/user/:id", tdSrvcs.UpdateUser)
	router.DELETE("/user/:id", tdSrvcs.DeleteUSer)

	err = router.Run(":8080")
	if err != nil {
		return
	}
}
