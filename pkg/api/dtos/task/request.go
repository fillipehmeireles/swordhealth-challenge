package dtos

type CreateTaskDto struct {
	Title        string `json:"title"`
	Summary      string `json:"summary"`
	TechnicianID int    `json:"technician_id"`
}

type UpdateTaskDto struct {
	Title        string `json:"title,omitempty"`
	Summary      string `json:"summary,omitempty"`
	TechnicianID int    `json:"technician_id,omitempty"`
}

type TaskStatusDto struct {
	Status bool `json:"status"`
}
