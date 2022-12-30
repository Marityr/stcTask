package service

import (
	"server/pkg/repository"
	"server/shemes"
)

type DataListService struct {
	repo repository.DataList
}

func NewDataListService(repo repository.DataList) *DataListService {
	return &DataListService{repo: repo}
}

func (s *DataListService) GetKey(num int) ([]shemes.QuizesAnswer, error) {
	return s.repo.GetKey(num)
}
