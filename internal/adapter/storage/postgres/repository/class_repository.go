package repository

import (
	"errors"

	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/storage/postgres"
	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/storage/postgres/repository/mapper"
	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/storage/postgres/repository/model"
	"github.com/KavindaKAAL/school-management-system-v2/internal/core/domain"
	"gorm.io/gorm"
)

type ClassRepository struct {
	database postgres.Database
}

func NewClassRepository(db postgres.Database) *ClassRepository {
	return &ClassRepository{
		db,
	}
}

var ErrTeacherAlreadyAssigned = errors.New("teacher is already assigned in the class")
var ErrTeacherNotAssigned = errors.New("teacher is not assigned in the class")

func (s *ClassRepository) GetAllClasses() ([]*domain.Class, error) {
	var classesPs []*model.ClassModel

	err := s.database.GetInstance().Preload("Students").Preload("Teacher").Find(&classesPs).Error
	if err != nil {
		return nil, err
	}
	classesDomain := mapper.ToDomainClassList(classesPs)
	return classesDomain, nil
}

func (s *ClassRepository) GetClassByName(className string) (*domain.Class, error) {
	var classPs *model.ClassModel

	err := s.database.GetInstance().Preload("Students").First(&classPs, "name = ?", className).Error
	if err != nil {
		return nil, err
	}

	classDomain := mapper.ToDomainClass(classPs)
	return classDomain, nil
}

func (s *ClassRepository) CreateClass(class *domain.Class) (*domain.Class, error) {

	classPs := mapper.FromDomainClass(class)

	err := s.database.GetInstance().Create(&classPs).Error
	if err != nil {
		return nil, err
	}

	return class, nil
}

func (s *ClassRepository) UpdateClass(class *domain.Class) (*domain.Class, error) {
	var updateClassPs model.ClassModel
	err := s.database.GetInstance().First(&updateClassPs, "name = ?", class.Name).Error

	if err != nil {
		return nil, err
	}
	updateClassPs.Subject = class.Subject

	err2 := s.database.GetInstance().Save(&updateClassPs).Error
	if err2 != nil {
		return nil, err2
	}

	classDomain := mapper.ToDomainClass(&updateClassPs)
	return classDomain, nil
}

func (s *ClassRepository) DeleteClassByName(className string) (bool, error) {
	var classPs model.ClassModel
	err := s.database.GetInstance().First(&classPs, "name = ?", className).Error

	if err != nil {
		return false, err
	}

	err2 := s.database.GetInstance().Unscoped().Delete(&classPs).Error

	if err2 != nil {
		return false, err2
	}

	return true, nil
}

func (s *ClassRepository) AssignTeacher(className string, teacherEmail string) error {

	return s.database.GetInstance().Transaction(func(tx *gorm.DB) error {
		var classPs model.ClassModel

		if err := tx.First(&classPs, "name = ?", className).Error; err != nil {
			return err
		}

		var teacherPs model.TeacherModel

		if err := tx.First(&teacherPs, "email = ?", teacherEmail).Error; err != nil {
			return err
		}

		// checking the teacher is already assigned.
		// Need to check another teacher is already assigned to the class.
		count := tx.Model(&classPs).Where("email = ?", teacherEmail).Association("Teacher").Count()
		if count > 0 {
			return ErrTeacherAlreadyAssigned
		}

		if err := tx.Model(&classPs).Association("Teacher").Append(&teacherPs); err != nil {
			return err
		}

		return nil
	})

}

func (s *ClassRepository) UnAssignTeacher(className string, teacherEmail string) error {

	return s.database.GetInstance().Transaction(func(tx *gorm.DB) error {
		var classPs model.ClassModel

		if err := tx.Preload("Teacher").First(&classPs, "name = ?", className).Error; err != nil {
			return err
		}

		var teacherPs model.TeacherModel

		if err := tx.First(&teacherPs, "email = ?", teacherEmail).Error; err != nil {
			return err
		}

		count := tx.Model(&classPs).Where("email = ?", teacherEmail).Association("Teacher").Count()
		if count == 0 {
			return ErrTeacherNotAssigned
		}

		if err := tx.Model(&classPs).Association("Teacher").Delete(&teacherPs); err != nil {
			return err
		}

		return nil
	})

}
