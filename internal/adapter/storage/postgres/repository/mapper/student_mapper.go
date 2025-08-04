package mapper

import (
	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/storage/postgres/repository/model"
	"github.com/KavindaKAAL/school-management-system-v2/internal/core/domain"
)

func ToDomainStudent(ps *model.StudentModel) *domain.Student {
	if ps == nil {
		return nil
	}

	return &domain.Student{
		Name:    ps.Name,
		Email:   ps.Email,
		Classes: ToDomainClassList(ps.Classes),
	}
}

func FromDomainStudent(ds *domain.Student) *model.StudentModel {
	if ds == nil {
		return nil
	}

	return &model.StudentModel{
		Name:  ds.Name,
		Email: ds.Email,
	}
}

func ToDomainStudentList(psList []*model.StudentModel) []*domain.Student {
	if psList == nil {
		return nil
	}

	dsList := make([]*domain.Student, len(psList))
	for i, s := range psList {
		dsList[i] = ToDomainStudent(s)
	}
	return dsList

}
