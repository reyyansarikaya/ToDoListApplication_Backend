package repository

import (
	"bootcamp/model"
)

type IRepository interface {
	Get() []model.ToDoModel
	Save(todo string)
}

func NewRepository() IRepository {
	return &Repository{
		Todolist: map[int]*model.ToDoModel{
			0: {
				ID:   0,
				Todo: "Dommy todo", // for pact test

			},
		},
	}
}

type Repository struct {
	Todolist map[int]*model.ToDoModel
}

func (r *Repository) Get() []model.ToDoModel {
	var todolist []model.ToDoModel
	for _, todo := range r.Todolist {
		todolist = append(todolist, *todo)
	}
	return todolist
}

func (r *Repository) Save(todo string) {
	todolist := r.Todolist
	lastID := 0
	if len(todolist) != 0 {
		lastID = len(todolist)
	}
	r.Todolist[lastID] = &model.ToDoModel{ID: lastID, Todo: todo}
}
