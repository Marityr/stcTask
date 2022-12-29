package service

import (
	"stcTask/server/pkg/repository"
	"stcTask/server/shemes"
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
