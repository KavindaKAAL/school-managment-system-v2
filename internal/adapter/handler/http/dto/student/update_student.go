package student

import (
	"fmt"

	"github.com/KavindaKAAL/school-management-system-v2/internal/core/port"
	"github.com/go-playground/validator"
)

type UpdateStudentDto struct {
	Name  string `json:"name" validate:"required,max=200"`
	Email string `json:"email" validate:"required,email"`
}

func EmptyUpdateStudentDto() port.Dto[UpdateStudentDto] {

	return &UpdateStudentDto{}
}

func (d *UpdateStudentDto) GetValue() *UpdateStudentDto {
	return d
}

func (d *UpdateStudentDto) ValidateErrors(errs validator.ValidationErrors) ([]string, error) {
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
