package usecase

import (
	"github.com/Mo3g4u/go-ddd-sample/domain/model"
	"github.com/Mo3g4u/go-ddd-sample/domain/repository"
	"github.com/Mo3g4u/go-ddd-sample/domain/service"
	"github.com/Mo3g4u/go-ddd-sample/interfaces/request"
)

type UserUsecase struct {
	r  repository.Repository
	us service.UserService
}

func NewUserUsecase(r repository.Repository, us service.UserService) *UserUsecase {
	return &UserUsecase{
		r:  r,
		us: us,
	}
}

func (u *UserUsecase) Create(req request.UserCreateRequest) error {
	b, err := u.us.Exists(req.Name)
	if err != nil {
		return err
	}
	if b {
		// ユーザーが存在する場合は独自エラーを返す
		return &model.UserExistsError{}
	}
	conf := model.UserCreateConfig{
		Name: req.Name,
	}
	user, err := model.NewUser(conf)
	if err != nil {
		return err
	}
	if err := u.r.CreateUser(user); err != nil {
		return err
	}
	return nil
}
