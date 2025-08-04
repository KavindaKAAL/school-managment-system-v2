package class

import (
	"fmt"

	"github.com/KavindaKAAL/school-management-system-v2/internal/core/domain"
	"github.com/go-playground/validator/v10"
)

type GetClassResDto struct {
	Name          *string                 `json:"name" validate:"required,max=50"`
	Subject       *string                 `json:"subject" validate:"required,max=50"`
	Teacher_Email *string                 `json:"teacher_email,omitempty"`
	Students      []*StudentInClassResDto `json:"students,omitempty"`
}

func EmptyGetClassResDto() *GetClassResDto {

	return &GetClassResDto{}
}

func (d *GetClassResDto) GetValue() *GetClassResDto {
	return d
}

func (d *GetClassResDto) ValidateErrors(errs validator.ValidationErrors) ([]string, error) {
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

func NewClassResDto(class *domain.Class) *GetClassResDto {
	var teacher_mail *string
	if class.Teacher != nil {
		teacher_mail = &class.Teacher.Email
	}
	return &GetClassResDto{
		Name:          &class.Name,
		Subject:       &class.Subject,
		Teacher_Email: teacher_mail,
		Students:      NewStudentsListResDto(class.Students),
	}
}

func NewClassesListResDto(classList []*domain.Class) []*GetClassResDto {
	if classList == nil {
		return nil
	}

	formattedClassList := make([]*GetClassResDto, len(classList))
	for i, c := range classList {
		formattedClassList[i] = NewClassResDto(c)
	}
	return formattedClassList

}
