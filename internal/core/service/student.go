package service

import (
	"github.com/KavindaKAAL/school-management-system-v2/internal/core/domain"
	"github.com/KavindaKAAL/school-management-system-v2/internal/core/port"
)

type StudentService struct {
	port.BaseService
	repo port.StudentRepository
}

func NewStudentService(repo port.StudentRepository) StudentService {
	return StudentService{
		BaseService: NewBaseService(),
		repo:        repo,
	}
}

func (s *StudentService) GetAllStudentService() ([]*domain.Student, error) {
	return s.repo.GetAllStudent()
}

func (s *StudentService) GetStudentByEmailService(email string) (*domain.Student, error) {
	return s.repo.GetStudentByEmail(email)
}

func (s *StudentService) CreateStudentService(student *domain.Student) error {
	return s.repo.CreateStudent(student)
}

func (s *StudentService) UpdateStudentService(student *domain.Student) error {
	return s.repo.UpdateStudent(student)
}

func (s *StudentService) DeleteStudentByEmailService(email string) error {
	return s.repo.DeleteStudentByEmail(email)
}

func (s *StudentService) EnrollStudentService(studentEmail string, className string) error {
	return s.repo.EnrollStudent(studentEmail, className)
}

func (s *StudentService) UnEnrollStudentService(studentEmail string, className string) error {
	return s.repo.UnEnrollStudent(studentEmail, className)

}
