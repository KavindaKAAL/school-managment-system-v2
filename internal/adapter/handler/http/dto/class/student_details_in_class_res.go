package class

import (
	"fmt"

	"github.com/KavindaKAAL/school-management-system-v2/internal/core/domain"
	"github.com/go-playground/validator/v10"
)

type StudentInClassResDto struct {
	Name  string `json:"student_name" validate:"required"`
	Email string `json:"student_email" validate:"required,email"`
}

func EmptyStudentInClassResDto() *StudentInClassResDto {

	return &StudentInClassResDto{}
}

func (d *StudentInClassResDto) GetValue() *StudentInClassResDto {
	return d
}

func (d *StudentInClassResDto) ValidateErrors(errs validator.ValidationErrors) ([]string, error) {
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

func NewStudentResDto(student *domain.Student) *StudentInClassResDto {

	return &StudentInClassResDto{
		Name:  student.Name,
		Email: student.Email,
	}
}

func NewStudentsListResDto(studentList []*domain.Student) []*StudentInClassResDto {
	if studentList == nil {
		return nil
	}

	formattedStudentList := make([]*StudentInClassResDto, len(studentList))
	for i, s := range studentList {
		formattedStudentList[i] = NewStudentResDto(s)
	}
	return formattedStudentList

}
