package models

type AlertUsers struct {
	BaseModel

	Name       string `json:"name"`
	Department string `json:"department"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
}

type AlertUsersRepo interface {
}
