package model

import (
	"gorm.io/gorm"
	"time"
)

func (User) TableName() string {return "users"}

const (
	ADMIN = "ADMIN"
	USER = "USER"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name" gorm:"not null" validate:"required"`
	Username  string `json:"username" gorm:"uniqueIndex,not null" validate:"required"`
	Password  []byte `gorm:"not null"`
	Role  string `json:"role" gorm:"not null" validate:"oneof=ADMIN USER,required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type RegisterRequest struct {
	Name      string `url:"name" json:"name" validate:"required"`
	Username  string `url:"username" json:"username" validate:"required"`
	Password  string `url:"password" json:"password" validate:"required"`
	Role  string `url:"role" json:"role" validate:"required,oneof=ADMIN USER"`
}

type LoginRequest struct {
	Username  string `url:"username" json:"username" validate:"required"`
	Password  string `url:"password" json:"password" validate:"required"`
}

type LoginResponse struct {
	TokenType  string `url:"tokenType" json:"tokenType" validate:"required"`
	Token  string `url:"token" json:"token" validate:"required"`
	Role  string `url:"role" json:"role" validate:"required"`
}