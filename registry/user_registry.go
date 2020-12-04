package registry

import (
	controller "golang-clean-architecture/interface/controller/user"
	presenter "golang-clean-architecture/interface/presenter/user"
	repository "golang-clean-architecture/interface/repository/user"
	interactor "golang-clean-architecture/usecase/interactor/user"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserRepository() repository.UserRepository {
	return repository.NewUserRepository(r.db)
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserPresenter() presenter.UserPresenter {
	return presenter.NewUserPresenter()
}
