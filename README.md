# School Management System API

A Go-based RESTful API for managing students, classes and teachers in a school. This system supports enrollment, assignment management, student,class and teacher CRUD operations.

---

## Project Structure

```
school-management-system/
├── cmd/
│   └── http/                      # Application entry point (main.go for HTTP server)
│       └── main.go
├── internal/
│   ├── adapter/                   # Adapters to external systems
│   │   ├── handler/               # Delivery layer (HTTP handlers using Gin)
│   │   │   └── http/
│   │   │       ├── student_handler.go
│   │   │       ├── class_handler.go
│   │   │       └── teacher_handler.go
│   │   ├── repository/            # Persistence layer
│   │      └── postgres/
│   │          ├── models/        # GORM models
│   │          │   ├── student_model.go
│   │          │   ├── class_model.go
│   │          │   └── teacher_model.go
│   │          ├── mappers/       # Mappers: convert between domain <-> DB
│   │          │   ├── student_mapper.go
│   │          │   ├── class_mapper.go
│   │          │   └── teacher_mapper.go
│   │          ├── student_repository.go
│   │          ├── class_repository.go
│   │          └── teacher_repository.go
│   │                  
│   │       
│   ├── core/
│      ├── domain/                # Enterprise entities (pure business models)
│      │   ├── student.go
│      │   ├── class.go
│      │   └── teacher.go
│      ├── port/                  # Interfaces
│      │   ├── repository/
│      │   │   ├── student_repository.go
│      │   │   ├── class_repository.go
│      │   │   └── teacher_repository.go
│      │   └── service/
│      │       ├── student_service.go
│      │       ├── class_service.go
│      │       └── teacher_service.go
│      ├── service/               # Application business logic
│          ├── student_service.go
│          ├── class_service.go
│          └── teacher_service.go
│                        
│   
├── scripts/                   
├── go.mod
├── go.sum
└── README.md

```

---

## API Endpoints

### Student Routes

| Method | Path                         | Description                        |
|--------|------------------------------|------------------------------------|
| POST   | `/api/v1/students`           | Create a new student               |
| GET    | `/api/v1/students`           | Get all students                   |
| GET    | `/api/v1/students/:email`       | Get a specific student by Email       |
| PUT    | `/api/v1/students`           | Update student by Email (in body)     |
| DELETE | `/api/v1/students/:email`       | Delete student by Email               |
| PUT    | `/api/v1/students/enroll`    | Enroll a student into a class      |
| PUT    | `/api/v1/students/unenroll`  | Unenroll a student from a class    |

### Class Routes

| Method | Path                 | Description                 |
|--------|----------------------|-----------------------------|
| POST   | `/api/v1/classes`    | Create a new class          |
| GET    | `/api/v1/classes`    | Get all classes             |
| GET    | `/api/v1/classes/:name` | Get a specific class by Name |
| DELETE | `/api/v1/classes/:name` | Delete a class             |
| PUT    | `/api/v1/classes/assignTeacher`    | Assign a teacher into a class      |
| PUT    | `/api/v1/classes/unAssignTeacher`  | Un-assign a teacher from a class    |

### Teacher Routes

| Method | Path                         | Description                        |
|--------|------------------------------|------------------------------------|
| POST   | `/api/v1/teachers`           | Create a new teacher               |
| GET    | `/api/v1/teachers`           | Get all teachers                   |
| GET    | `/api/v1/teachers/:email`       | Get a specific teacher by Email      |
| PUT    | `/api/v1/teachers`           | Update teacher by Email (in body)     |
| DELETE | `/api/v1/teachers/:email`       | Delete teacher by Email               |


---
