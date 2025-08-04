package student

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type TeacherInClassInStudentResDto struct {
	Name  string `json:"teacher_name,omitempty"`
	Email string `json:"teacher_email,omitempty"`
}

func EmptyTeacherInClassInStudentResDto() *TeacherInClassInStudentResDto {

	return &TeacherInClassInStudentResDto{}
}

func (d *TeacherInClassInStudentResDto) GetValue() *TeacherInClassInStudentResDto {
	return d
}

func (d *TeacherInClassInStudentResDto) ValidateErrors(errs validator.ValidationErrors) ([]string, error) {
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
