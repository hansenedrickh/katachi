package handler

import "github.com/hansenedrickh/katachi/usecase"

type Handler struct {
	usecase usecase.Usecase
}

func NewHandler(u usecase.Usecase) Handler {
	return Handler{
		usecase: u,
	}
}
