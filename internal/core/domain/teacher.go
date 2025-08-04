package domain

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Teacher struct {
	ID        uint64
	Name      string `json:"name" validate:"required,max=50"`
	Email     string `json:"email" validate:"required,email"`
	Classes   []*Class
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTeacher(name string, email string) (*Teacher, error) {

	u := Teacher{
		Name:  name,
		Email: email,
	}

	if err := u.Validate(); err != nil {
		return nil, err
	}
	return &u, nil
}

func (teacher *Teacher) GetValue() *Teacher {
	return teacher
}

func (teacher *Teacher) Validate() error {
	validate := validator.New()
	return validate.Struct(teacher)
}
