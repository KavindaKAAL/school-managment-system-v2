package teacher

import (
	"fmt"

	"github.com/KavindaKAAL/school-management-system-v2/internal/core/port"
	"github.com/go-playground/validator"
)

type UpdateTeacherDto struct {
	Name  string `json:"name" validate:"required,max=200"`
	Email string `json:"email" validate:"required,email"`
}

func EmptyUpdateTeacherDto() port.Dto[UpdateTeacherDto] {

	return &UpdateTeacherDto{}
}

func (d *UpdateTeacherDto) GetValue() *UpdateTeacherDto {
	return d
}

func (d *UpdateTeacherDto) ValidateErrors(errs validator.ValidationErrors) ([]string, error) {
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
