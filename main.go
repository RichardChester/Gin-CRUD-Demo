package main

import (
	"github.com/gin-gonic/gin"
	"go-crud-demo/controllers"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		userController := new(controllers.UserController)
		v1.POST("/users",userController.Create)
		v1.GET("/users",userController.GetAll)
		v1.GET("/users/:id",userController.GetOne)
		v1.DELETE("/users/:id",userController.DeleteOne)
		v1.PUT("/users/:id",userController.UpdateOne)
	}

	router.NoRoute(func(context *gin.Context) {
		context.String(404,"fuck you")
	})

	router.Run()
}


