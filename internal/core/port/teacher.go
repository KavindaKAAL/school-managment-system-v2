package port

import (
	"github.com/KavindaKAAL/school-management-system-v2/internal/core/domain"
)

type TeacherRepository interface {
	GetAllTeachers() ([]*domain.Teacher, error)
	GetTeacherByEmail(email string) (*domain.Teacher, error)
	CreateTeacher(teacher *domain.Teacher) error
	UpdateTeacher(teacher *domain.Teacher) (*domain.Teacher, error)
	DeleteTeacherByEmail(email string) (bool, error)
}

type TeacherService interface {
	GetAllTeacherService() ([]*domain.Teacher, error)
	GetTeacherByEmailService(email string) (*domain.Teacher, error)
	CreateTeacherService(teacher *domain.Teacher) error
	UpdateTeacherService(teacher *domain.Teacher) (*domain.Teacher, error)
	DeleteTeacherByEmailService(email string) (bool, error)
}
