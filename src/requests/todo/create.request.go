package requests

type CreateTodoRequest struct {
	Name string `json:"name" binding:"required"`
}
