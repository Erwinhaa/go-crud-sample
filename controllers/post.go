package controllers

import (
	"net/http"

	"myapp/models"
	"myapp/services"

	"github.com/gin-gonic/gin"
)

type PostController struct{}

func (controller *PostController) GetPost(ctx *gin.Context) {
	s := services.GetService()

	posts, err := s.GetPosts(ctx)

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, models.PostsPublicReturn{
		Status: true,
		Data:   posts,
	})
}

func (controller *PostController) CreatePost(ctx *gin.Context) {
	var input models.Post

	s := services.GetTransaction()

	err := ctx.ShouldBind(&input)

	if err != nil {
		panic(err)
	}

	post, _ := s.CreatePost(ctx.Request.Context(), input)

	s.Commit()

	ctx.JSON(http.StatusOK, models.PostPublicReturn{
		Status: true,
		Data:   post,
	})
}
