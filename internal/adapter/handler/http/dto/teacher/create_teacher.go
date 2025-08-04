package teacher

import (
	"fmt"

	"github.com/KavindaKAAL/school-management-system-v2/internal/core/port"
	"github.com/go-playground/validator"
)

type CreateTeacherDto struct {
	Name  string `json:"name" binding:"required" validate:"required,max=200"`
	Email string `json:"email" binding:"required" validate:"required"`
}

func EmptyCreateTeacherDto() port.Dto[CreateTeacherDto] {

	return &CreateTeacherDto{}
}

func (d *CreateTeacherDto) GetValue() *CreateTeacherDto {
	return d
}

func (d *CreateTeacherDto) ValidateErrors(errs validator.ValidationErrors) ([]string, error) {
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
