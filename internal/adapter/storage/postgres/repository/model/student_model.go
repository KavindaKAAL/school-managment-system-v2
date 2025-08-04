package model

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type StudentModel struct {
	gorm.Model
	Name    string        `json:"name" validate:"required,max=50"`
	Email   string        `json:"email" gorm:"unique" validate:"required,email"`
	Classes []*ClassModel `json:"classes" gorm:"many2many:student_classes;constraint:OnDelete:CASCADE;"`
}

func NewStudent(name string, email string) (*StudentModel, error) {

	u := StudentModel{
		Name:  name,
		Email: email,
	}

	if err := u.Validate(); err != nil {
		return nil, err
	}
	return &u, nil
}

func (student *StudentModel) GetValue() *StudentModel {
	return student
}

func (student *StudentModel) Validate() error {
	validate := validator.New()
	return validate.Struct(student)
}
