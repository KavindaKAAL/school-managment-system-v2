package repository

import (
	"errors"

	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/storage/postgres"
	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/storage/postgres/repository/mapper"
	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/storage/postgres/repository/model"
	"github.com/KavindaKAAL/school-management-system-v2/internal/core/domain"
	"gorm.io/gorm"
)

var ErrStudentAlreadyEnrolled = errors.New("student is already enrolled in the class")
var ErrStudentNotEnrolled = errors.New("student is not enrolled in the class")
var ErrEmailAlreadyInUse = errors.New("email is already use")
var ErrUserNotFound = errors.New("user not found")

type StudentRepository struct {
	database postgres.Database
}

func NewStudentRepository(db postgres.Database) *StudentRepository {
	return &StudentRepository{
		db,
	}
}

func (s *StudentRepository) GetAllStudent() ([]*domain.Student, error) {
	var studentsPostgres []*model.StudentModel

	err := s.database.GetInstance().Preload("Classes.Teacher").Find(&studentsPostgres).Error
	if err != nil {
		return nil, err
	}

	studentsDomain := mapper.ToDomainStudentList(studentsPostgres)
	return studentsDomain, nil
}

func (s *StudentRepository) GetStudentByEmail(email string) (*domain.Student, error) {
	var studentPs *model.StudentModel

	err := s.database.GetInstance().Preload("Classes.Teacher").First(&studentPs, "email = ?", email).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrUserNotFound
		} else {
			return nil, err
		}
	}

	studentDomain := mapper.ToDomainStudent(studentPs)
	return studentDomain, nil
}

func (s *StudentRepository) CreateStudent(student *domain.Student) error {

	var existingStudentPs *model.StudentModel
	if err := s.database.GetInstance().First(&existingStudentPs, "email = ?", student.Email).Error; err == nil {
		return ErrEmailAlreadyInUse
	} else if err != gorm.ErrRecordNotFound {
		return err
	}

	var existingTeacherPs *model.TeacherModel
	if err := s.database.GetInstance().First(&existingTeacherPs, "email = ?", student.Email).Error; err == nil {
		return ErrEmailAlreadyInUse
	} else if err != gorm.ErrRecordNotFound {
		return err
	}

	studentPs := mapper.FromDomainStudent(student)
	err := s.database.GetInstance().Create(&studentPs).Error
	return err
}

func (s *StudentRepository) UpdateStudent(student *domain.Student) (*domain.Student, error) {

	var updateStudentPs *model.StudentModel
	err := s.database.GetInstance().First(&updateStudentPs, "email = ?", student.Email).Error

	if err != nil {
		return nil, err
	}
	updateStudentPs.Name = student.Name

	err2 := s.database.GetInstance().Save(&updateStudentPs).Error
	if err2 != nil {
		return nil, err2
	}

	studentDomain := mapper.ToDomainStudent(updateStudentPs)
	return studentDomain, nil
}

func (s *StudentRepository) DeleteStudentByEmail(email string) error {

	var studentPs model.StudentModel

	if err := s.database.GetInstance().First(&studentPs, "email = ?", email).Error; err == nil {
		return s.database.GetInstance().Unscoped().Delete(&studentPs).Error
	} else {
		if err == gorm.ErrRecordNotFound {
			return ErrUserNotFound
		} else {
			return err
		}
	}
}

func (s *StudentRepository) EnrollStudent(studentEmail string, className string) error {

	return s.database.GetInstance().Transaction(func(tx *gorm.DB) error {
		var studentPs model.StudentModel

		if err := tx.Preload("Classes").First(&studentPs, "email = ?", studentEmail).Error; err != nil {
			return err
		}

		var classPs model.ClassModel

		if err := tx.First(&classPs, "name = ?", className).Error; err != nil {
			return err
		}

		count := tx.Model(&studentPs).Where("name = ?", classPs.Name).Association("Classes").Count()
		if count > 0 {
			return ErrStudentAlreadyEnrolled
		}

		if err := tx.Model(&studentPs).Association("Classes").Append(&classPs); err != nil {
			return err
		}

		return nil
	})

}

func (s *StudentRepository) UnEnrollStudent(studentEmail string, className string) error {

	return s.database.GetInstance().Transaction(func(tx *gorm.DB) error {
		var studentPs model.StudentModel

		if err := tx.Preload("Classes").First(&studentPs, "email = ?", studentEmail).Error; err != nil {
			return err
		}

		var classPs model.ClassModel

		if err := tx.First(&classPs, "name = ?", className).Error; err != nil {
			return err
		}

		count := tx.Model(&studentPs).Where("name = ?", className).Association("Classes").Count()
		if count == 0 {
			return ErrStudentNotEnrolled
		}

		if err := tx.Model(&studentPs).Association("Classes").Delete(&classPs); err != nil {
			return err
		}

		return nil
	})

}
