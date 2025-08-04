package model

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type TeacherModel struct {
	gorm.Model
	Name    string        `json:"name" validate:"required,max=50"`
	Email   string        `json:"email" gorm:"unique" validate:"required,email"`
	Classes []*ClassModel `gorm:"foreignKey:TeacherID;constraint:OnDelete:CASCADE;"`
}

func NewTeacher(name string, email string) (*TeacherModel, error) {

	u := TeacherModel{
		Name:  name,
		Email: email,
	}

	if err := u.Validate(); err != nil {
		return nil, err
	}
	return &u, nil
}

func (teacher *TeacherModel) GetValue() *TeacherModel {
	return teacher
}

func (teacher *TeacherModel) Validate() error {
	validate := validator.New()
	return validate.Struct(teacher)
}
