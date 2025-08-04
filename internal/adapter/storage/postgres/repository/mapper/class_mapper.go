package mapper

import (
	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/storage/postgres/repository/model"
	"github.com/KavindaKAAL/school-management-system-v2/internal/core/domain"
)

func ToDomainClass(pc *model.ClassModel) *domain.Class {
	if pc == nil {
		return nil
	}
	students := ToDomainStudentList(pc.Students)
	teacher := ToDomainTeacher(pc.Teacher)
	return &domain.Class{
		Name:     pc.Name,
		Subject:  pc.Subject,
		Students: students,
		Teacher:  teacher,
	}
}

func FromDomainClass(dc *domain.Class) *model.ClassModel {
	if dc == nil {
		return nil
	}

	return &model.ClassModel{
		Name:    dc.Name,
		Subject: dc.Subject,
	}
}

func ToDomainClassList(pcList []*model.ClassModel) []*domain.Class {
	if pcList == nil {
		return nil
	}

	dsList := make([]*domain.Class, len(pcList))
	for i, c := range pcList {
		dsList[i] = ToDomainClass(c)
	}
	return dsList

}
