package domain

type User struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDTO struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u User) ToDTO() *UserDTO {
	var dto UserDTO
	dto.UUID = u.UUID
	dto.Email = u.Email
	dto.Name = u.Name

	return &dto
}
