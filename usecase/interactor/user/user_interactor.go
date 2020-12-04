package interactor

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	model "golang-clean-architecture/domain/model/user"
	"golang-clean-architecture/helper"
	presenter "golang-clean-architecture/interface/presenter/user"
	repository "golang-clean-architecture/interface/repository/user"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type userInteractor struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

type UserInteractor interface {
	Register(user model.RegisterRequest) (*model.User, error)
	Login(user model.LoginRequest) (*model.LoginResponse, error)
}

func NewUserInteractor(r repository.UserRepository, p presenter.UserPresenter) UserInteractor {
	return userInteractor{
		UserRepository: r,
		UserPresenter:  p,
	}
}

func (ui userInteractor) Register(user model.RegisterRequest) (*model.User, error) {
	if err := ui.UserRepository.CheckUsernameExist(user.Username); err != nil {
		return nil, err
	}

	pw, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, err
	}
	fmt.Println(user.Password)
	fmt.Println(pw)
	data := model.User{
		Name:     user.Name,
		Username: user.Username,
		Password: pw,
		Role:     user.Role,
	}
	res, err := ui.UserRepository.Register(data)
	if err != nil {
		return nil, err
	}
	return ui.UserPresenter.ResponseRegister(*res), nil
}

func (ui userInteractor) Login(user model.LoginRequest) (*model.LoginResponse, error) {
	res, err := ui.UserRepository.Login(user.Username, []byte(user.Password))
	if err != nil {
		return nil, err
	}

	token := jwt.New(jwt.SigningMethodHS512)
	claim := token.Claims.(jwt.MapClaims)
	claim["id"] = res.ID
	claim["username"] = res.Username
	claim["role"] = res.Role
	claim["exp"] = time.Now().Add(time.Minute * helper.JWT_EXPIRY_MINUTES).Unix()

	tokenresult, err := token.SignedString([]byte(helper.JWT_SECURITY_TOKEN))
	if err != nil {
		return nil, err
	}

	return ui.UserPresenter.ResponseLogin(*res, tokenresult), nil
}
