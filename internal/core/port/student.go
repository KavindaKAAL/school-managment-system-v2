package port

import (
	"github.com/KavindaKAAL/school-management-system-v2/internal/core/domain"
)

type StudentRepository interface {
	GetAllStudent() ([]*domain.Student, error)
	GetStudentByEmail(email string) (*domain.Student, error)
	CreateStudent(student *domain.Student) error
	UpdateStudent(student *domain.Student) (*domain.Student, error)
	DeleteStudentByEmail(email string) error
	EnrollStudent(studentEmail string, className string) error
	UnEnrollStudent(studentEmail string, className string) error
}

type StudentService interface {
	GetAllStudentService() ([]*domain.Student, error)
	GetStudentByEmailService(email string) (*domain.Student, error)
	CreateStudentService(student *domain.Student) error
	UpdateStudentService(student *domain.Student) (*domain.Student, error)
	DeleteStudentByEmailService(email string) error
	EnrollStudentService(studentEmail string, className string) error
	UnEnrollStudentService(studentEmail string, className string) error
}
