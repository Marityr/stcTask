package service

import (
	"stcTask/server/pkg/repository"
	"stcTask/server/shemes"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type DataList interface {
	GetKey(num int) ([]shemes.QuizesAnswer, error)
}

type Service struct {
	DataList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		DataList: NewDataListService(repos.DataList),
	}
}
