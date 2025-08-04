package port

import (
	"github.com/KavindaKAAL/school-management-system-v2/internal/core/domain"
)

type ClassRepository interface {
	GetAllClasses() ([]*domain.Class, error)
	GetClassByName(name string) (*domain.Class, error)
	CreateClass(class *domain.Class) (*domain.Class, error)
	UpdateClass(class *domain.Class) (*domain.Class, error)
	DeleteClassByName(name string) (bool, error)
	AssignTeacher(className string, teacherEmail string) error
	UnAssignTeacher(className string, teacherEmail string) error
}

type ClassService interface {
	GetAllClassService() ([]*domain.Class, error)
	GetClassByNameService(name string) (*domain.Class, error)
	CreateClassService(class *domain.Class) (*domain.Class, error)
	UpdateClassService(class *domain.Class) (*domain.Class, error)
	DeleteClassByNameService(name string) (bool, error)
	AssignTeacherService(className string, teacherEmail string) error
	UnAssignTeacherService(className string, teacherEmail string) error
}
