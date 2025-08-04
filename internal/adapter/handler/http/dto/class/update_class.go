package class

import (
	"fmt"

	"github.com/KavindaKAAL/school-management-system-v2/internal/core/port"
	"github.com/go-playground/validator"
)

type UpdateClassDto struct {
	Name    string `json:"name" validate:"required,max=50"`
	Subject string `json:"subject" validate:"required,max=50"`
}

func EmptyUpdateClassDto() port.Dto[UpdateClassDto] {

	return &UpdateClassDto{}
}

func (d *UpdateClassDto) GetValue() *UpdateClassDto {
	return d
}

func (d *UpdateClassDto) ValidateErrors(errs validator.ValidationErrors) ([]string, error) {
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
