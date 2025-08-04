package domain

import (
	"fmt"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

var namePattern = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)

type Class struct {
	ID        uint64
	Name      string `json:"name" validate:"required"`
	Subject   string `json:"subject" validate:"required"`
	Students  []*Student
	TeacherID *uint    `json:"teacher_id"`
	Teacher   *Teacher `json:"teacher"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewClass(name string, subject string) (*Class, error) {

	u := Class{
		Name:    name,
		Subject: subject,
	}

	if err := u.Validate(); err != nil {
		return nil, err
	}
	return &u, nil
}

func (class *Class) GetValue() *Class {
	return class
}

func (class *Class) Validate() error {
	validate := validator.New()

	if !namePattern.MatchString(class.Name) {
		return fmt.Errorf("name can only contain letters, numbers, underscores, and hyphens without spaces")
	}
	return validate.Struct(class)
}
