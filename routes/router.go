package routes

import (
	"myapp/controllers"
	"myapp/middlewares"

	"github.com/gin-gonic/gin"
)

func WebRouter(router *gin.Engine) {
	userController := controllers.UserController{}
	postController := new(controllers.PostController)

	router.POST("/login", userController.Login)

	authRoute := router.Group("")
	authRoute.Use(middlewares.IsUser())
	userRoute := authRoute.Group("user")
	{
		userRoute.POST("/", userController.CreateUser)
		userRoute.GET("/", userController.GetUsers)
		userRoute.GET("/:id", userController.GetUserByID)
		userRoute.PUT("/:id", userController.UpdateUser)
		userRoute.DELETE("/:id", userController.DeleteUser)
		userRoute.GET("/me", userController.GetMe)
		userRoute.GET("/me/posts", userController.GetMyPost)
	}

	postRoute := authRoute.Group("post")
	{
		postRoute.GET("/", postController.GetPost)
		postRoute.POST("/", postController.CreatePost)
		postRoute.PUT("/:id", postController.UpdatePost)
		postRoute.DELETE("/:id", postController.DeletePost)
	}
}
