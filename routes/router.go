package routes

import (
	"myapp/controllers"

	"github.com/gin-gonic/gin"
)

func WebRouter(router *gin.Engine) {
	postController := new(controllers.PostController)

	postRoute := router.Group("post")
	{
		postRoute.GET("/", postController.GetPost)
		postRoute.POST("/", postController.CreatePost)
	}
}
