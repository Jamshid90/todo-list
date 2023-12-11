package models

type Task struct {
	UUID        string `json:"uuid"`
	Title       string `json:"title"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt    string `json:"update_at"`
}

type AddTaskRequest struct {
	Title       string `json:"title"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
	Description string `json:"description"`
}

type AddTaskResponse struct {
	Task Task `json:"task"`
}

type GetTaskResponse struct {
	Task Task `json:"task"`
}

type ListTaskResponse struct {
	List []Task `json:"list"`
}
