package class

import (
	"fmt"

	"github.com/KavindaKAAL/school-management-system-v2/internal/core/port"
	"github.com/go-playground/validator"
)

type CreateClassDto struct {
	Name    string `json:"name" binding:"required" validate:"required,max=50"`
	Subject string `json:"subject" binding:"required" validate:"required,,max=50"`
}

func EmptyCreateClassDto() port.Dto[CreateClassDto] {

	return &CreateClassDto{}
}

func (d *CreateClassDto) GetValue() *CreateClassDto {
	return d
}

func (d *CreateClassDto) ValidateErrors(errs validator.ValidationErrors) ([]string, error) {
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
