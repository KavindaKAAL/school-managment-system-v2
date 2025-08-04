package student

import (
	"fmt"

	"github.com/KavindaKAAL/school-management-system-v2/internal/core/domain"
	"github.com/go-playground/validator/v10"
)

type ClassInStudentResDto struct {
	ClassName   string `json:"class_name" validate:"required"`
	SubjectName string `json:"subject" validate:"required"`
}

func EmptyClassInStudentResDto() *ClassInStudentResDto {

	return &ClassInStudentResDto{}
}

func (d *ClassInStudentResDto) GetValue() *ClassInStudentResDto {
	return d
}

func (d *ClassInStudentResDto) ValidateErrors(errs validator.ValidationErrors) ([]string, error) {
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

func NewClassResDto(class *domain.Class) *ClassInStudentResDto {

	return &ClassInStudentResDto{
		ClassName:   class.Name,
		SubjectName: class.Subject,
	}
}

func NewClassesListResDto(classList []*domain.Class) []*ClassInStudentResDto {
	if classList == nil {
		return nil
	}

	formattedStudentList := make([]*ClassInStudentResDto, len(classList))
	for i, s := range classList {
		formattedStudentList[i] = NewClassResDto(s)
	}
	return formattedStudentList

}
