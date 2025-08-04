package service

import (
	"github.com/KavindaKAAL/school-management-system-v2/internal/core/domain"
	"github.com/KavindaKAAL/school-management-system-v2/internal/core/port"
)

type TeacherService struct {
	port.BaseService
	repo port.TeacherRepository
}

func NewTeacherService(repo port.TeacherRepository) TeacherService {
	return TeacherService{
		BaseService: NewBaseService(),
		repo:        repo,
	}
}

func (s *TeacherService) GetAllTeacherService() ([]*domain.Teacher, error) {
	teachers, err := s.repo.GetAllTeachers()
	return teachers, err
}

func (s *TeacherService) GetTeacherByEmailService(email string) (*domain.Teacher, error) {
	teacher, err := s.repo.GetTeacherByEmail(email)
	return teacher, err
}

func (s *TeacherService) CreateTeacherService(teacher *domain.Teacher) error {
	err := s.repo.CreateTeacher(teacher)
	return err
}

func (s *TeacherService) UpdateTeacherService(teacher *domain.Teacher) error {
	return s.repo.UpdateTeacher(teacher)
}

func (s *TeacherService) DeleteTeacherByEmailService(teacherEmail string) error {
	return s.repo.DeleteTeacherByEmail(teacherEmail)
}
