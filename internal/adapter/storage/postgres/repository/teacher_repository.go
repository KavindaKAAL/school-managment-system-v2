package repository

import (
	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/storage/postgres"
	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/storage/postgres/repository/mapper"
	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/storage/postgres/repository/model"
	"github.com/KavindaKAAL/school-management-system-v2/internal/core/domain"
	"gorm.io/gorm"
)

type TeacherRepository struct {
	database postgres.Database
}

func NewTeacherRepository(db postgres.Database) *TeacherRepository {
	return &TeacherRepository{
		db,
	}
}

func (s *TeacherRepository) GetAllTeachers() ([]*domain.Teacher, error) {
	var teachersPs []*model.TeacherModel

	err := s.database.GetInstance().Preload("Classes").Find(&teachersPs).Error
	if err != nil {
		return nil, err
	}
	teachersDomain := mapper.ToDomainTeacherList(teachersPs)
	return teachersDomain, nil
}

func (s *TeacherRepository) GetTeacherByEmail(teacherEmail string) (*domain.Teacher, error) {
	var teacherPs *model.TeacherModel

	err := s.database.GetInstance().Preload("Classes").First(&teacherPs, "email = ?", teacherEmail).Error
	if err != nil {
		return nil, err
	}
	teacherDomain := mapper.ToDomainTeacher(teacherPs)
	return teacherDomain, nil
}

func (s *TeacherRepository) CreateTeacher(teacher *domain.Teacher) error {

	var existingTeacherPs *model.TeacherModel
	if err := s.database.GetInstance().First(&existingTeacherPs, "email = ?", teacher.Email).Error; err == nil {
		return ErrEmailAlreadyInUse
	} else if err != gorm.ErrRecordNotFound {
		return err
	}

	var existingStudentPs *model.StudentModel
	if err := s.database.GetInstance().First(&existingStudentPs, "email = ?", teacher.Email).Error; err == nil {
		return ErrEmailAlreadyInUse
	} else if err != gorm.ErrRecordNotFound {
		return err
	}

	teacherPs := mapper.FromDomainTeacher(teacher)

	err := s.database.GetInstance().Create(&teacherPs).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *TeacherRepository) UpdateTeacher(teacher *domain.Teacher) (*domain.Teacher, error) {
	var updateTeacherPs model.TeacherModel
	err := s.database.GetInstance().First(&updateTeacherPs, "email = ?", teacher.Email).Error

	if err != nil {
		return nil, err
	}
	updateTeacherPs.Name = teacher.Name
	err2 := s.database.GetInstance().Save(&updateTeacherPs).Error
	if err2 != nil {
		return nil, err2
	}

	teacherDomain := mapper.ToDomainTeacher(&updateTeacherPs)
	return teacherDomain, nil
}

func (s *TeacherRepository) DeleteTeacherByEmail(teacherEmail string) (bool, error) {
	var teacherPs model.TeacherModel
	err := s.database.GetInstance().First(&teacherPs, "email = ?", teacherEmail).Error

	if err != nil {
		return false, err
	}

	err2 := s.database.GetInstance().Unscoped().Delete(&teacherPs).Error

	if err2 != nil {
		return false, err2
	}

	return true, nil
}
