package service

import (
	"github.com/KavindaKAAL/school-management-system-v2/internal/core/domain"
	"github.com/KavindaKAAL/school-management-system-v2/internal/core/port"
)

type ClassService struct {
	port.BaseService
	repo port.ClassRepository
}

func NewClassService(repo port.ClassRepository) ClassService {
	return ClassService{
		BaseService: NewBaseService(),
		repo:        repo,
	}
}

func (s *ClassService) GetAllClassService() ([]*domain.Class, error) {
	classes, err := s.repo.GetAllClasses()
	return classes, err
}

func (s *ClassService) GetClassByNameService(name string) (*domain.Class, error) {
	class, err := s.repo.GetClassByName(name)
	return class, err
}

func (s *ClassService) CreateClassService(class *domain.Class) error {
	return s.repo.CreateClass(class)

}

func (s *ClassService) UpdateClassService(class *domain.Class) error {

	return s.repo.UpdateClass(class)
}

func (s *ClassService) DeleteClassByNameService(className string) error {

	return s.repo.DeleteClassByName(className)
}

func (s *ClassService) AssignTeacherService(className string, teacherEmail string) error {
	err := s.repo.AssignTeacher(className, teacherEmail)
	return err
}

func (s *ClassService) UnAssignTeacherService(className string, teacherEmail string) error {
	err := s.repo.UnAssignTeacher(className, teacherEmail)
	return err
}
