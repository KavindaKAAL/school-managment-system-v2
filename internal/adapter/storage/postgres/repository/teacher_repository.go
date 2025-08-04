package repository

import (
	"errors"

	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/storage/postgres"
	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/storage/postgres/repository/mapper"
	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/storage/postgres/repository/model"
	"github.com/KavindaKAAL/school-management-system-v2/internal/core/domain"
	"gorm.io/gorm"
)

var ErrUserAssignedToSomeClasses = errors.New("user still assigned to classes")

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

func (s *TeacherRepository) GetTeacherByEmail(email string) (*domain.Teacher, error) {
	var teacherPs *model.TeacherModel

	err := s.database.GetInstance().Preload("Classes").First(&teacherPs, "email = ?", email).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrUserNotFound
		} else {
			return nil, err
		}
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

func (s *TeacherRepository) UpdateTeacher(teacher *domain.Teacher) error {
	var updateTeacherPs model.TeacherModel
	err := s.database.GetInstance().First(&updateTeacherPs, "email = ?", teacher.Email).Error

	if err != nil {
		return err
	}
	updateTeacherPs.Name = teacher.Name
	return s.database.GetInstance().Save(&updateTeacherPs).Error

}

func (s *TeacherRepository) DeleteTeacherByEmail(email string) error {
	var teacherPs model.TeacherModel

	if err := s.database.GetInstance().Preload("Classes").First(&teacherPs, "email = ?", email).Error; err == nil {
		count := len(teacherPs.Classes)
		if count > 0 {
			return ErrUserAssignedToSomeClasses
		}
		return s.database.GetInstance().Unscoped().Delete(&teacherPs).Error
	} else {
		if err == gorm.ErrRecordNotFound {
			return ErrUserNotFound
		} else {
			return err
		}
	}

}
