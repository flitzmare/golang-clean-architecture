package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	model "golang-clean-architecture/domain/model/user"
	interactor "golang-clean-architecture/usecase/interactor/user"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

type UserController interface {
	Register(context *gin.Context)
	Login(context *gin.Context)
}

func NewUserController(ui interactor.UserInteractor) UserController {
	return &userController{userInteractor:ui}
}

func (c userController) Register(context *gin.Context){
	user := model.RegisterRequest{
		Name:     context.PostForm("name"),
		Username: context.PostForm("username"),
		Password: context.PostForm("password"),
		Role: context.PostForm("role"),
	}

	v := validator.New()
	err := v.Struct(user)
	if err != nil {
		context.JSON(422, err.Error())
		return
	}

	_, err = c.userInteractor.Register(user)
	if err != nil {
		context.JSON(400, err.Error())
		return
	}

	context.JSON(200, "OK")
}

func (c userController) Login(context *gin.Context){
	user := model.LoginRequest{
		Username: context.PostForm("username"),
		Password: context.PostForm("password"),
	}

	v := validator.New()
	err := v.Struct(user)
	if err != nil {
		context.JSON(422, err.Error())
		return
	}

	res, err := c.userInteractor.Login(user)
	if err != nil {
		context.JSON(400, err.Error())
		return
	}

	context.JSON(200, res)
}
