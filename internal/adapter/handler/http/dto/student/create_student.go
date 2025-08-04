package student

import (
	"fmt"

	"github.com/KavindaKAAL/school-management-system-v2/internal/core/port"
	"github.com/go-playground/validator"
)

type CreateStudentDto struct {
	Name  string `json:"name" binding:"required" validate:"required,max=50"`
	Email string `json:"email" binding:"required" validate:"required,email"`
}

func EmptyCreateStudentDto() port.Dto[CreateStudentDto] {

	return &CreateStudentDto{}
}

func (d *CreateStudentDto) GetValue() *CreateStudentDto {
	return d
}

func (d *CreateStudentDto) ValidateErrors(errs validator.ValidationErrors) ([]string, error) {
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
