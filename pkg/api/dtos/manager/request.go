package dtos

type CreateManagerDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type UpdateManagerDTO struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty"`
}
