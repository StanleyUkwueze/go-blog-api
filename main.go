package main

import (
	loadinitializers "apicrud/LoadInitializers"
	"apicrud/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	loadinitializers.LoadEnvVariables()
	loadinitializers.ConnectToDb()
}
func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)

	r.GET("/posts", controllers.PostFetch)

	r.GET("/posts/:id", controllers.PostFetchById)

	r.PUT("/posts/:id", controllers.PostUpdate)

	r.DELETE("/posts/:id", controllers.PostDelete)
	r.Run()
}
