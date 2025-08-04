package student

import (
	"fmt"

	"github.com/KavindaKAAL/school-management-system-v2/internal/core/port"
	"github.com/go-playground/validator"
)

type GetStudentReqDto struct {
	Email string `uri:"email" binding:"required" validate:"required"`
}

func EmptyGetStudentReqDto() port.Dto[GetStudentReqDto] {

	return &GetStudentReqDto{}
}

func (d *GetStudentReqDto) GetValue() *GetStudentReqDto {
	return d
}

func (d *GetStudentReqDto) ValidateErrors(errs validator.ValidationErrors) ([]string, error) {
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
