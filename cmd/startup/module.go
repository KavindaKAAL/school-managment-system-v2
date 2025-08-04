package startup

import (
	"context"

	"github.com/KavindaKAAL/school-management-system-v2/config"
	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/handler/http/controller"
	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/storage/postgres"
	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/storage/postgres/repository"
	"github.com/KavindaKAAL/school-management-system-v2/internal/core/port"
	"github.com/KavindaKAAL/school-management-system-v2/internal/core/service"
)

type Module port.Module[module]

type module struct {
	Context        context.Context
	Env            *config.Env
	DB             postgres.Database
	StudentService service.StudentService
	ClassService   service.ClassService
	TeacherService service.TeacherService
}

func (m *module) GetInstance() *module {
	return m
}

func (m *module) Controllers() []port.Controller {
	return []port.Controller{
		controller.NewStudentController(&m.StudentService),
		controller.NewClassController(&m.ClassService),
		controller.NewTeacherController(&m.TeacherService),
	}
}

func NewModule(context context.Context, env *config.Env, db postgres.Database) Module {
	studentRepo := repository.NewStudentRepository(db)
	classRepo := repository.NewClassRepository(db)
	teacherRepo := repository.NewTeacherRepository(db)

	studentService := service.NewStudentService(studentRepo)
	classService := service.NewClassService(classRepo)
	teacherService := service.NewTeacherService(teacherRepo)

	return &module{
		Context:        context,
		Env:            env,
		DB:             db,
		StudentService: studentService,
		ClassService:   classService,
		TeacherService: teacherService,
	}
}
