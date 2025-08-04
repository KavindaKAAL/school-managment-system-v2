package teacher

import (
	"fmt"

	"github.com/KavindaKAAL/school-management-system-v2/internal/core/domain"
	"github.com/go-playground/validator/v10"
)

type ClassInTeacherResDto struct {
	ClassName   string `json:"class_name" validate:"required"`
	SubjectName string `json:"subject" validate:"required"`
}

func EmptyClassInTeacherResDto() *ClassInTeacherResDto {

	return &ClassInTeacherResDto{}
}

func (d *ClassInTeacherResDto) GetValue() *ClassInTeacherResDto {
	return d
}

func (d *ClassInTeacherResDto) ValidateErrors(errs validator.ValidationErrors) ([]string, error) {
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

func NewClassResDto(class *domain.Class) *ClassInTeacherResDto {

	return &ClassInTeacherResDto{
		ClassName:   class.Name,
		SubjectName: class.Subject,
	}
}

func NewClassesListResDto(classList []*domain.Class) []*ClassInTeacherResDto {
	if classList == nil {
		return nil
	}

	formattedTeacherList := make([]*ClassInTeacherResDto, len(classList))
	for i, s := range classList {
		formattedTeacherList[i] = NewClassResDto(s)
	}
	return formattedTeacherList

}
