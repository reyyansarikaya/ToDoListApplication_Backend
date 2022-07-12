package service

import (
	"bootcamp/model"
	"bootcamp/repository"
	"net/http"
)

type IService interface {
	Get() []model.ToDoModel
	Save(todo string) int
}

func NewService(repo repository.IRepository) *DefaultToDoService {
	return &DefaultToDoService{repo: repo}
}

type DefaultToDoService struct {
	repo repository.IRepository
}

func (s *DefaultToDoService) Get() []model.ToDoModel {
	return s.repo.Get()
}

func (s *DefaultToDoService) Save(todo string) int {
	s.repo.Save(todo)
	return http.StatusOK
}
