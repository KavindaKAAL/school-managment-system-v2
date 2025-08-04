package mapper

import (
	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/storage/postgres/repository/model"
	"github.com/KavindaKAAL/school-management-system-v2/internal/core/domain"
)

func ToDomainTeacher(tp *model.TeacherModel) *domain.Teacher {
	if tp == nil {
		return nil
	}

	return &domain.Teacher{
		Name:    tp.Name,
		Email:   tp.Email,
		Classes: ToDomainClassList(tp.Classes),
	}
}

func FromDomainTeacher(td *domain.Teacher) *model.TeacherModel {
	if td == nil {
		return nil
	}

	return &model.TeacherModel{
		Name:  td.Name,
		Email: td.Email,
	}
}

func ToDomainTeacherList(tpList []*model.TeacherModel) []*domain.Teacher {
	if tpList == nil {
		return nil
	}

	dsList := make([]*domain.Teacher, len(tpList))
	for i, t := range tpList {
		dsList[i] = ToDomainTeacher(t)
	}
	return dsList

}
