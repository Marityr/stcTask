package service

import "stcTask/server/pkg/repository"

type DataListService struct {
	repo repository.DataList
}

func NewDataListService(repo repository.DataList) *DataListService {
	return &DataListService{repo: repo}
}

func (s *DataListService) GetKey(key string) (string, error) {
	return s.repo.GetKey(key)
}
