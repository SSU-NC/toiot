package handler

import (
	"github.com/KumKeeHyun/toiot/application/usecase"
)

type Handler struct {
	ru usecase.RegistUsecase
	eu usecase.EventUsecase
}

func NewHandler(ru usecase.RegistUsecase, eu usecase.EventUsecase) *Handler {
	return &Handler{
		ru: ru,
		eu: eu,
	}
}
