package services

import "github.com/iqbaludinm/library-api/repositories"

type BaseService struct {
	repo repositories.RepoInterface
}

type ServiceInterface interface {
	BookService
}

func NewService(repo repositories.RepoInterface) ServiceInterface {
	return &BaseService{repo: repo}
}


