package service

import (
	"awesomeProject1/src/domain/entities"
	"fmt"
)

type RepositoryUser interface {
	List() ([]entities.User, error)
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
