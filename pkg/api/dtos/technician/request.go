package dtos

type CreateTechnicianDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type UpdateTechnicianDto struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty"`
}
