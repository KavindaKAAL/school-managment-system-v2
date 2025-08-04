package student

import (
	"fmt"

	"github.com/KavindaKAAL/school-management-system-v2/internal/core/domain"
	"github.com/go-playground/validator/v10"
)

type GetStudentResDto struct {
	Name    *string                 `json:"name" validate:"required,max=200"`
	Email   *string                 `json:"email" validate:"required,email"`
	Classes []*ClassInStudentResDto `json:"classes,omitempty"`
}

func EmptyGetStudentResDto() *GetStudentResDto {

	return &GetStudentResDto{}
}

func (d *GetStudentResDto) GetValue() *GetStudentResDto {
	return d
}

func (d *GetStudentResDto) ValidateErrors(errs validator.ValidationErrors) ([]string, error) {
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

func NewStudentResDto(student *domain.Student) *GetStudentResDto {

	return &GetStudentResDto{
		Name:    &student.Name,
		Email:   &student.Email,
		Classes: NewClassesListResDto(student.Classes),
	}
}

func NewStudentsListResDto(studentList []*domain.Student) []*GetStudentResDto {
	if studentList == nil {
		return nil
	}

	formattedStudentList := make([]*GetStudentResDto, len(studentList))
	for i, s := range studentList {
		formattedStudentList[i] = NewStudentResDto(s)
	}
	return formattedStudentList

}
