package repository

import (
	"errors"
	"fmt"
	model "golang-clean-architecture/domain/model/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepository struct {
	*gorm.DB
}

type UserRepository interface {
	CheckUsernameExist(username string) error
	Register(user model.User) (*model.User, error)
	Login(username string, password []byte) (*model.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (u userRepository) CheckUsernameExist(username string) error {
	var count int64
	u.Model(&model.User{}).Where("username = ?", username).Count(&count)
	if count >= 1 {
		return errors.New("This username is already exist")
	}
	return nil
}

func (u userRepository) Register(user model.User) (*model.User, error) {
	tx := u.Create(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (u userRepository) Login(username string, password []byte) (*model.User, error) {
	user := model.User{}
	tx := u.Where("username = ?", username).First(&user)
	if tx.Error == gorm.ErrRecordNotFound {
		return nil, tx.Error
	}
	if tx.Error != nil {
		return nil, tx.Error
	}
	fmt.Println(user.Password)
	fmt.Println(password)
	if err := bcrypt.CompareHashAndPassword(user.Password, password); err != nil {
		return nil, err
	}
	return &user, nil
}
