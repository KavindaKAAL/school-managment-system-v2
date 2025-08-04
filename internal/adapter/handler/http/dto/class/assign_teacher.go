package class

import (
	"fmt"

	"github.com/KavindaKAAL/school-management-system-v2/internal/core/port"
	"github.com/go-playground/validator"
)

type AssignTeacherToClassDto struct {
	TeacherEmail string `json:"teacher_email" binding:"required" validate:"required,email"`
	ClassName    string `json:"class_name" binding:"required" validate:"required"`
}

func EmptyAssignTeacherToClassDto() port.Dto[AssignTeacherToClassDto] {

	return &AssignTeacherToClassDto{}
}

func (d *AssignTeacherToClassDto) GetValue() *AssignTeacherToClassDto {
	return d
}

func (d *AssignTeacherToClassDto) ValidateErrors(errs validator.ValidationErrors) ([]string, error) {
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
