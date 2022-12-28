package repositories

import (
	"golang-todo/src/models"

	"gorm.io/gorm"
)

type TodoRepo interface {
	GetAll() ([]models.Todo, error)
	Get(string) (models.Todo, error)
	Create(string) (models.Todo, error)
	Update(string, string) (models.Todo, error)
	Delete(string) (bool, error)
	ToggleDone(string) (models.Todo, error)
	ChangeAllDone(bool) (bool, error)
}

type todoRepo struct {
	DB *gorm.DB
}

func NewTodoRepo(db *gorm.DB) TodoRepo {
	return &todoRepo{DB: db}
}

func (todoRepo *todoRepo) GetAll() ([]models.Todo, error) {
	var todos []models.Todo

	result := todoRepo.DB.Find(&todos)

	if result.Error != nil {
		return nil, result.Error
	}

	return todos, nil

}

func (todoRepo *todoRepo) Get(id string) (models.Todo, error) {

	var todo models.Todo

	result := todoRepo.DB.Where("id = ?", id).First(&todo)

	if result.Error != nil {
		return models.Todo{}, result.Error
	}

	return todo, nil

}

func (todoRepo *todoRepo) Create(name string) (models.Todo, error) {
	todo := models.Todo{
		Name: name,
		Done: false,
	}

	result := todoRepo.DB.Model(models.Todo{}).Create(&todo)

	if result.Error != nil {
		return models.Todo{}, nil
	}

	return todo, nil

}

func (todoRepo *todoRepo) Update(id string, name string) (models.Todo, error) {
	todo := models.Todo{
		Name: name,
		Done: false,
	}

	result := todoRepo.DB.Model(models.Todo{}).Where("id = ?", id).Updates(&todo)

	if result.Error != nil {
		return models.Todo{}, nil
	}

	return todo, nil

}

func (todoRepo *todoRepo) Delete(id string) (bool, error) {

	result := todoRepo.DB.Model(models.Todo{}).Delete(&models.Todo{}, id)

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil

}

func (todoRepo *todoRepo) ToggleDone(id string) (models.Todo, error) {

	var todo models.Todo

	result1 := todoRepo.DB.Where("id = ?", id).First(&todo)

	if result1.Error != nil {
		return models.Todo{}, result1.Error
	}

	if todo.Done {
		todo.Done = false
	} else {
		todo.Done = true
	}

	result2 := todoRepo.DB.Model(models.Todo{}).Where("id = ?", id).Updates(&todo)

	if result2.Error != nil {
		return models.Todo{}, result1.Error
	}

	return todo, nil
}

func (todoRepo *todoRepo) ChangeAllDone(done bool) (bool, error) {
	result := todoRepo.DB.Exec("UPDATE todos SET done = ?", done)

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil

}
