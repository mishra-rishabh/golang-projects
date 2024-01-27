package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mishra-rishabh/json_crud_api/controllers"
	"github.com/mishra-rishabh/json_crud_api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	fmt.Println("hey bro!")

	/* Router */
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.FetchAllPosts)
	r.GET("/posts/:id", controllers.FetchPost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)
	r.Run()
}
