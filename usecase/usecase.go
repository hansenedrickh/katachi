package usecase

import "github.com/hansenedrickh/katachi/repository"

type Usecase interface {

}

type usecase struct {
	repository repository.Repository
}

func NewUsecase(r repository.Repository) Usecase {
	return &usecase{
		repository: r,
	}
}
