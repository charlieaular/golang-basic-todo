package requests

type MarkAllDone struct {
	Done *bool `json:"done" binding:"required"`
}
