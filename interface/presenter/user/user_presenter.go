package presenter

import (
	model "golang-clean-architecture/domain/model/user"
	"golang-clean-architecture/helper"
)

type userPresenter struct {
}

type UserPresenter interface {
	ResponseRegister(user model.User) *model.User
	ResponseLogin(user model.User, token string) *model.LoginResponse
}

func NewUserPresenter() UserPresenter {
	return &userPresenter{}
}

func (p userPresenter) ResponseRegister(user model.User) *model.User {
	return &user
}

func (p userPresenter) ResponseLogin(user model.User, token string) *model.LoginResponse {
	return &model.LoginResponse{
		TokenType: helper.DEFAULT_TOKEN_TYPE,
		Token:     token,
		Role:      user.Role,
	}
}
