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

func (controller *PostController) UpdatePost(ctx *gin.Context) {
	var input models.PostInput

	ctx.ShouldBind(&input)
	ctx.ShouldBindUri(&input)

	s := services.GetTransaction()

	post, err := s.UpdatePost(ctx, input)

	s.Commit()

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, models.PostPublicReturn{
		Status: true,
		Data: &models.PostReturn{
			ID:    post.ID,
			Title: post.Title,
			Body:  post.Body,
		},
	})
}

func (controllers *PostController) DeletePost(ctx *gin.Context) {
	var post models.Post
	s := services.GetTransaction()

	ctx.ShouldBindUri(&post)

	message, err := s.DeletePost(ctx, &post)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	s.Commit()

	ctx.JSON(http.StatusOK, message)
}
