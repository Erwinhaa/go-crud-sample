package controllers

import (
	"myapp/models"
	"myapp/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (controller *UserController) CreateUser(ctx *gin.Context) {
	var userInput models.UserCreateInput
	s := services.GetTransaction()

	if err := ctx.ShouldBind(&userInput); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.UserPublicReturn{
			Status: false,
			Data:   nil,
		})
		return
	}

	user, _ := s.CreateUser(ctx, &userInput)

	s.Commit()

	ctx.JSON(http.StatusOK, models.UserPublicReturn{
		Status: true,
		Data:   user,
	})
}

func (controller *UserController) GetUsers(ctx *gin.Context) {
	s := services.GetService()

	users, _ := s.GetUsers(ctx.Request.Context())

	ctx.JSON(http.StatusOK, models.UsersPublicReturn{
		Status: true,
		Data:   users,
	})
}

func (controller *UserController) GetUserByID(ctx *gin.Context) {
	var user *models.User

	if err := ctx.ShouldBindUri(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	s := services.GetService()

	foundUser, _ := s.GetUserByID(ctx.Request.Context(), user)

	ctx.JSON(http.StatusOK, models.UserPublicReturn{
		Status: true,
		Data:   foundUser,
	})
}

func (controller *UserController) UpdateUser(ctx *gin.Context) {
	var user *models.UserUpdateInput

	if err := ctx.ShouldBindUri(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	s := services.GetTransaction()

	updatedUser, _ := s.UpdateUser(ctx.Request.Context(), user)

	s.Commit()

	ctx.JSON(http.StatusOK, models.UserPublicReturn{
		Status: true,
		Data:   updatedUser,
	})
}

func (controller *UserController) DeleteUser(ctx *gin.Context) {
	var user *models.User

	if err := ctx.ShouldBindUri(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	s := services.GetTransaction()

	message, _ := s.DeleteUser(ctx.Request.Context(), user)

	s.Commit()

	ctx.JSON(http.StatusOK, message)
}

func (controller *UserController) Login(ctx *gin.Context) {
	var user *models.UserLoginInput

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	s := services.GetService()

	token, _ := s.Login(ctx.Request.Context(), user)

	ctx.JSON(http.StatusOK, models.UserTokenReturn{
		Status: true,
		Token:  token,
	})
}

func (controller *UserController) GetMe(ctx *gin.Context) {

	s := services.GetService()

	user, _ := s.GetMe(ctx.Request.Context())

	ctx.JSON(http.StatusOK, models.UserPublicReturn{
		Status: true,
		Data:   user,
	})
}

func (controller *UserController) GetMyPost(ctx *gin.Context) {
	s := services.GetService()

	posts, _ := s.GetMyPost(ctx.Request.Context())

	ctx.JSON(http.StatusOK, posts)
}
