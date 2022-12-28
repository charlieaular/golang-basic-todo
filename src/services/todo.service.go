package services

import (
	"golang-todo/src/models"
	"golang-todo/src/repositories"
)

type TodoService interface {
	GetAll() ([]models.Todo, error)
	Get(string) (models.Todo, error)
	Create(string) (models.Todo, error)
	Update(string, string) (models.Todo, error)
	Delete(string) (bool, error)
	ToggleDone(string) (models.Todo, error)
	ChangeAllDone(bool) (bool, error)
}

type todoService struct {
	todoRepo repositories.TodoRepo
}

func NewTodoService(todoRepo repositories.TodoRepo) TodoService {
	return &todoService{todoRepo: todoRepo}
}

func (todoService *todoService) GetAll() ([]models.Todo, error) {

	return todoService.todoRepo.GetAll()
}

func (todoService *todoService) Create(name string) (models.Todo, error) {
	return todoService.todoRepo.Create(name)
}

func (todoService *todoService) Get(id string) (models.Todo, error) {
	return todoService.todoRepo.Get(id)
}

func (todoService *todoService) Update(id string, name string) (models.Todo, error) {
	return todoService.todoRepo.Update(id, name)
}

func (todoService *todoService) Delete(id string) (bool, error) {
	return todoService.todoRepo.Delete(id)
}

func (todoService *todoService) ToggleDone(id string) (models.Todo, error) {
	return todoService.todoRepo.ToggleDone(id)
}

func (todoService *todoService) ChangeAllDone(done bool) (bool, error) {
	return todoService.todoRepo.ChangeAllDone(done)
}
