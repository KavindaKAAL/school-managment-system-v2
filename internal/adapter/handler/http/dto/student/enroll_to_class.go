package student

import (
	"fmt"

	"github.com/KavindaKAAL/school-management-system-v2/internal/core/port"
	"github.com/go-playground/validator"
)

type EnrollStudentToClassDto struct {
	StudentEmail string `json:"student_email" binding:"required" validate:"required,email"`
	ClassName    string `json:"class_name" binding:"required" validate:"required"`
}

func EmptyEnrollStudentToClassDto() port.Dto[EnrollStudentToClassDto] {

	return &EnrollStudentToClassDto{}
}

func (d *EnrollStudentToClassDto) GetValue() *EnrollStudentToClassDto {
	return d
}

func (d *EnrollStudentToClassDto) ValidateErrors(errs validator.ValidationErrors) ([]string, error) {
	var msgs []string
	for _, err := range errs {
		switch err.Tag() {
		case "required":
			msgs = append(msgs, fmt.Sprintf("%s is required", err.Field()))

		default:
			msgs = append(msgs, fmt.Sprintf("%s is invalid", err.Field()))
		}
	}
	return msgs, nil
}
