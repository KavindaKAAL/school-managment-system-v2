package teacher

import (
	"fmt"

	"github.com/go-playground/validator"
)

type GetTeacherReqDto struct {
	Email string `uri:"email" binding:"required" validate:"required"`
}

func EmptyGetTeacherReqDto() *GetTeacherReqDto {

	return &GetTeacherReqDto{}
}

func (d *GetTeacherReqDto) GetValue() *GetTeacherReqDto {
	return d
}

func (d *GetTeacherReqDto) ValidateErrors(errs validator.ValidationErrors) ([]string, error) {
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
