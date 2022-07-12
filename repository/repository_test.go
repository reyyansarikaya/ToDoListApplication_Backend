//go:build unit_test
// +build unit_test

package repository_test

import (
	"bootcamp/model"
	"bootcamp/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultToDoRepository_Get(t *testing.T) {
	todolist := map[int]*model.ToDoModel{
		0: {
			ID:   0,
			Todo: "Dommy todo",
		},
	}

	r := repository.NewRepository()
	var restodo []model.ToDoModel
	for _, todo := range todolist {
		restodo = append(restodo, *todo)
	}
	result := r.Get()

	assert.ElementsMatch(t, restodo, result)
}

func TestRepository_Save(t *testing.T) {
	todolist := map[int]*model.ToDoModel{
		0: {
			ID:   0,
			Todo: "buy some milk",
		},
		1: {
			ID:   1,
			Todo: "go to swim",
		},
	}

	r := repository.NewRepository()
	todo := "buy some chocolate"
	lastID := 1
	r.Save(todo)

	assert.Contains(t, todolist, lastID)
	assert.Equal(t, "buy some chocolate", todo)
	assert.Equal(t, 1, lastID)
}
