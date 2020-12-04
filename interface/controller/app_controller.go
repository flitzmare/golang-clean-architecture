package controller

import controller "golang-clean-architecture/interface/controller/user"

type AppController struct {
	User interface{ controller.UserController }
}
