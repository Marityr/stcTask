package service

import "stcTask/server/pkg/repository"

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type DataList interface {
	GetKey(key string) (string, error)
}

type Service struct {
	DataList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		DataList: NewDataListService(repos.DataList),
	}
}
