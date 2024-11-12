package service

import (
	"github.com/heartlezz7/go-echo-todolist/model"
	"github.com/heartlezz7/go-echo-todolist/repo"
	"github.com/rs/zerolog/log"
)

func CreateUserService(user model.CreateUser) error {

	err := repo.CreateUser(user)

	if err != nil {
		return err
	}
	return nil
}

func GetUserDetailService(id int, user *model.UserData) error {

	userId := repo.NewUserId(id)

	err := userId.GetUserDetail(user)
	if err != nil {
		log.Err(err).Msg("fail from service")
		return err
	}

	return nil
}

func LoginService(input model.UserLogin, userData model.UserLogin) error {

	return nil

}
