package class

import (
	"fmt"

	"github.com/KavindaKAAL/school-management-system-v2/internal/core/port"
	"github.com/go-playground/validator"
)

type GetClassReqDto struct {
	Name string `uri:"name" binding:"required" validate:"required"`
}

func EmptyGetClassReqDto() port.Dto[GetClassReqDto] {

	return &GetClassReqDto{}
}

func (d *GetClassReqDto) GetValue() *GetClassReqDto {
	return d
}

func (d *GetClassReqDto) ValidateErrors(errs validator.ValidationErrors) ([]string, error) {
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
