package models

type Task struct {
	ID        string `json:"id"`
	Assignee  string `json:"assignee"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	Deadline  string `json:"deadline"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateTask struct {
	Assignee  string `json:"assignee"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	Deadline  string `json:"deadline"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ListTasks struct {
	Tasks []Task `json:"tasks"`
}

type Overdue struct {
	Timed string `json:"timed"`
	Limit int64  `json:"limit"`
	Page  int64  `json:"page"`
}
