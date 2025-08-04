package model

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var namePattern = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)

type ClassModel struct {
	gorm.Model
	Name      string          `json:"name" gorm:"unique" validate:"required,max=50"`
	Subject   string          `json:"subject" validate:"required,max=50"`
	Students  []*StudentModel `json:"students" gorm:"many2many:student_classes;constraint:OnDelete:CASCADE;"`
	TeacherID *uint           `json:"teacher_id"`
	Teacher   *TeacherModel   `json:"teacher" gorm:"constraint:OnDelete:CASCADE;"`
}

func NewClass(name string, subject string) (*ClassModel, error) {

	u := ClassModel{
		Name:    name,
		Subject: subject,
	}

	if err := u.Validate(); err != nil {
		return nil, err
	}
	return &u, nil
}

func (class *ClassModel) GetValue() *ClassModel {
	return class
}

func (class *ClassModel) Validate() error {
	validate := validator.New()

	if !namePattern.MatchString(class.Name) {
		return fmt.Errorf("name can only contain letters, numbers, underscores, and hyphens without spaces")
	}
	return validate.Struct(class)
}
