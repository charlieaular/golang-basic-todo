package handlers

import (
	requests "golang-todo/src/requests/todo"
	"golang-todo/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todoHandler struct {
	todoService services.TodoService
}

func NewTodoHandler(todoService services.TodoService) todoHandler {
	return todoHandler{todoService: todoService}
}

func (h *todoHandler) GetAll(c *gin.Context) {

	todos, err := h.todoService.GetAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":    err.Error(),
			"status": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"todos":  todos,
	})
}

func (h *todoHandler) Get(c *gin.Context) {
	todoId := c.Param("id")
	todo, err := h.todoService.Get(todoId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":    err.Error(),
			"status": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"todo":   todo,
	})

}

func (h *todoHandler) Create(c *gin.Context) {

	var createTodoRequest requests.CreateTodoRequest

	err := c.ShouldBind(&createTodoRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":    err.Error(),
			"status": false,
		})
		return
	}

	todo, err := h.todoService.Create(createTodoRequest.Name)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":    err.Error(),
			"status": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"todo":   todo,
	})

}

func (h *todoHandler) Update(c *gin.Context) {

	var updateTodoRequest requests.UpdateTodoRequest

	err := c.ShouldBind(&updateTodoRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":    err.Error(),
			"status": false,
		})
		return
	}

	todoId := c.Param("id")

	todo, err := h.todoService.Update(todoId, updateTodoRequest.Name)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":    err.Error(),
			"status": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"todo":   todo,
	})

}

func (h *todoHandler) Delete(c *gin.Context) {

	todoId := c.Param("id")

	_, err := h.todoService.Delete(todoId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":    err.Error(),
			"status": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})

}

func (h *todoHandler) ToggleDone(c *gin.Context) {

	todoId := c.Param("id")

	todo, err := h.todoService.ToggleDone(todoId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":    err.Error(),
			"status": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"todo":   todo,
	})

}

func (h *todoHandler) ChangeAllDone(c *gin.Context) {

	var markAllDone requests.MarkAllDone

	err := c.ShouldBind(&markAllDone)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":    err.Error(),
			"status": false,
		})
		return
	}

	_, err = h.todoService.ChangeAllDone(*markAllDone.Done)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":    err.Error(),
			"status": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})

}
