//go:build unit_test
// +build unit_test

package service_test

import (
	"bootcamp/mock"
	"bootcamp/model"
	"bootcamp/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestDefaultToDoService_Get(t *testing.T) {

	repositoryReturn := []model.ToDoModel{
		0: {
			ID:   0,
			Todo: "buy some milk",
		},
		1: {
			ID:   1,
			Todo: "go to swim",
		},
	}

	exceptedTodo := make([]model.ToDoModel, 0)
	exceptedTodo = append(exceptedTodo, model.ToDoModel{
		ID:   0,
		Todo: "buy some milk",
	}, model.ToDoModel{
		ID:   1,
		Todo: "go to swim",
	})

	repository := mock.NewMockIRepository(gomock.NewController(t))
	repository.EXPECT().
		Get().
		Return(repositoryReturn).
		Times(1)

	serv := service.NewService(repository)
	actualTodo := serv.Get()

	assert.Equal(t, &exceptedTodo, &actualTodo)
}

func TestDefaultToDoService_Save(t *testing.T) {
	expect := http.StatusOK
	todo := "buy some milk"
	repository := mock.NewMockIRepository(gomock.NewController(t))
	repository.
		EXPECT().
		Save(todo).
		Return().
		Times(1)

	serv := service.NewService(repository)
	actualResult := serv.Save(todo)

	assert.Equal(t, &expect, &actualResult)
}
