package teacher

import (
	"fmt"

	"github.com/KavindaKAAL/school-management-system-v2/internal/core/domain"
	"github.com/go-playground/validator/v10"
)

type GetTeacherResDto struct {
	Name    *string                 `json:"name" validate:"required,max=200"`
	Email   *string                 `json:"email" validate:"required,email"`
	Classes []*ClassInTeacherResDto `json:"classes,omitempty"`
}

func EmptyGetTeacherResDto() *GetTeacherResDto {

	return &GetTeacherResDto{}
}

func (d *GetTeacherResDto) GetValue() *GetTeacherResDto {
	return d
}

func (d *GetTeacherResDto) ValidateErrors(errs validator.ValidationErrors) ([]string, error) {
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

func NewTeacherResDto(teacher *domain.Teacher) *GetTeacherResDto {

	return &GetTeacherResDto{
		Name:    &teacher.Name,
		Email:   &teacher.Email,
		Classes: NewClassesListResDto(teacher.Classes),
	}
}

func NewTeachersListResDto(teacherList []*domain.Teacher) []*GetTeacherResDto {
	if teacherList == nil {
		return nil
	}

	formattedTeacherList := make([]*GetTeacherResDto, len(teacherList))
	for i, s := range teacherList {
		formattedTeacherList[i] = NewTeacherResDto(s)
	}
	return formattedTeacherList

}
