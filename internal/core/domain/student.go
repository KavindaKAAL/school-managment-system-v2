package domain

import (
	"time"

	"github.com/go-playground/validator"
)

type Student struct {
	ID        uint64
	Name      string `validate:"required,max=50"`
	Email     string `validate:"required,email"`
	Classes   []*Class
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewStudent(name string, email string) (*Student, error) {

	u := Student{
		Name:  name,
		Email: email,
	}

	if err := u.Validate(); err != nil {
		return nil, err
	}
	return &u, nil
}

func (student *Student) GetValue() *Student {
	return student
}

func (student *Student) Validate() error {
	validate := validator.New()
	return validate.Struct(student)
}
