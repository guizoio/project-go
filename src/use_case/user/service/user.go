package service

import (
	"awesomeProject1/src/domain/entities"
	"awesomeProject1/src/use_case/user/request"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
)

type RepositoryUser interface {
	List() ([]entities.User, error)
	Create(data entities.User) error
}

type UserService struct {
	Repository RepositoryUser
}

func NewUserService(Repository RepositoryUser) UserService {
	return UserService{Repository}
}

func (r UserService) List() ([]entities.User, error) {
	list, err := r.Repository.List()
	fmt.Println(list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r UserService) Create(data request.User) error {
	uuidGenerator, err := uuid.NewV4()
	if err != nil {
		errorString := errors.New("Error uuid.New v4, error:" + err.Error())
		return errorString
	}
	user := entities.User{
		ID:       uuidGenerator.String(),
		Name:     data.Name,
		Login:    data.Login,
		Password: data.Password,
		Email:    data.Email,
	}
	return r.Repository.Create(user)
}
