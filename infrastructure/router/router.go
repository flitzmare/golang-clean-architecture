package router

import (
	"github.com/gin-gonic/gin"
	"golang-clean-architecture/interface/controller"
)

func NewRouter(c controller.AppController) *gin.Engine {
	r := gin.Default()
	r.POST("/register", c.User.Register)
	r.POST("/login", c.User.Login)
	return r
}
