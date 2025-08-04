package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/KavindaKAAL/school-management-system-v2/cmd/startup"
	"github.com/KavindaKAAL/school-management-system-v2/tests/utils"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

var createTemporaryStudent = func(t *testing.T) {

	client := resty.New()
	client.SetRetryCount(5)
	client.SetRetryWaitTime(2 * time.Second)

	body := map[string]interface{}{
		"name":  "Lahiru",
		"email": "lahiru@gmail.com",
	}

	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post("http://localhost:8081/api/v1/students")

	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
}

var createTemporaryClass = func(t *testing.T) {

	client := resty.New()
	client.SetRetryCount(5)
	client.SetRetryWaitTime(2 * time.Second)

	body := map[string]interface{}{
		"name":    "english-grade-12",
		"subject": "English",
	}

	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post("http://localhost:8081/api/v1/classes")

	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
}

var enrollTemporaryStudentToTemporaryClass = func(t *testing.T) {

	client := resty.New()
	client.SetRetryCount(5)
	client.SetRetryWaitTime(2 * time.Second)

	body := map[string]interface{}{
		"student_email": "lahiru@gmail.com",
		"class_name":    "english-grade-12",
	}

	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Put("http://localhost:8081/api/v1/students/enroll")

	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
}

func TestIntegrationCreatetStudent(t *testing.T) {
	router, module, teardown := startup.TestServer()
	go router.Start(module.GetInstance().Env.ServerHost, module.GetInstance().Env.ServerPort)
	time.Sleep(1 * time.Second)

	defer teardown()

	t.Run("It should return successfully created message when creating a new student successfully", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		body := map[string]interface{}{
			"name":  "Lahiru",
			"email": "lahiru@gmail.com",
		}

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Post("http://localhost:8081/api/v1/students")

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 201, resp.StatusCode())
		assert.Equal(t, "{\"message\":\"Successfully created\"}", resp.String())

	})

	t.Run("It should return error message when creating a new student using already existing email", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		body := map[string]interface{}{
			"name":  "Lahiru",
			"email": "lahiru@gmail.com",
		}

		// 1st rest call
		_, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Post("http://localhost:8081/api/v1/students")

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		body2 := map[string]interface{}{
			"name":  "Kavinda",
			"email": "lahiru@gmail.com",
		}
		// 2nd rest call
		resp2, err2 := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body2).
			Post("http://localhost:8081/api/v1/students")

		if err2 != nil {
			t.Fatalf("Request failed: %v", err2)
		}

		assert.Equal(t, 409, resp2.StatusCode())
		assert.Equal(t, "{\"error\":\"Email is already used\"}", resp2.String())

	})

	t.Run("It should return error message when creating a new student using missing required field email", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		body := map[string]interface{}{
			"name": "Lahiru",
			"age":  25,
		}

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Post("http://localhost:8081/api/v1/students")

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 400, resp.StatusCode())
		assert.Equal(t, "{\"error\":\"Invalid request body\"}", resp.String())

	})

	t.Run("It should return error message when creating a new student using missing required field name", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		body := map[string]interface{}{
			"email": "lahiru@gmail.com",
			"age":   25,
		}

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Post("http://localhost:8081/api/v1/students")

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 400, resp.StatusCode())
		assert.Equal(t, "{\"error\":\"Invalid request body\"}", resp.String())

	})

	t.Run("It should return error message when creating a new student using incorrect value for the name field", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		body := map[string]interface{}{
			"name":  25,
			"email": "lahiru@gmail.com",
		}

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Post("http://localhost:8081/api/v1/students")

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 400, resp.StatusCode())
		assert.Equal(t, "{\"error\":\"Invalid type for field name\"}", resp.String())

	})

	t.Run("It should return error message when creating a new student using incorrect value for the email field", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		body := map[string]interface{}{
			"name":  "Anshan",
			"email": 25,
		}

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Post("http://localhost:8081/api/v1/students")

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 400, resp.StatusCode())
		assert.Equal(t, "{\"error\":\"Invalid type for field email\"}", resp.String())

	})

	t.Run("It should return error message when creating a new student using incorrect request method", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		body := map[string]interface{}{
			"name":  "Anshan",
			"email": "lahiru@gmail.com",
		}

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Patch("http://localhost:8081/api/v1/students")

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 404, resp.StatusCode())
		assert.Equal(t, "{\"error\":\"Route not found\"}", resp.String())

	})

	t.Run("It should return error message when creating a new student using incorrect json format", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		invalidBody := `{
			name:  "Anshan",
			email: "lahiru@gmail.com",
		}`

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(invalidBody).
			Post("http://localhost:8081/api/v1/students")

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 400, resp.StatusCode())
		assert.Equal(t, "{\"error\":\"Malformed JSON: Syntax error\"}", resp.String())

	})

}

func TestIntegrationGetSpecificStudent(t *testing.T) {
	router, module, teardown := startup.TestServer()
	go router.Start(module.GetInstance().Env.ServerHost, module.GetInstance().Env.ServerPort)
	time.Sleep(1 * time.Second)

	defer teardown()

	// create a tempory student correctly
	createTemporaryStudent(t)

	t.Run("It should return a message contains the student details", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		param := "lahiru@gmail.com"

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(param).
			Get(fmt.Sprintf("http://localhost:8081/api/v1/students/%s", param))

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 200, resp.StatusCode())

	})

	t.Run("It should return an error message of student not found", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		param := "lahiru@gmail.com"

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(param).
			Get(fmt.Sprintf("http://localhost:8081/api/v1/students/%s", param))

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 404, resp.StatusCode())
		assert.Equal(t, "{\"error\":\"Student is not found\"}", resp.String())

	})
}

func TestIntegrationUpdateStudent(t *testing.T) {
	router, module, teardown := startup.TestServer()
	go router.Start(module.GetInstance().Env.ServerHost, module.GetInstance().Env.ServerPort)
	time.Sleep(1 * time.Second)

	defer teardown()

	// create a tempory student correctly
	createTemporaryStudent(t)

	t.Run("It should return a message contains successfully updated", func(t *testing.T) {
		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		body := map[string]interface{}{
			"name":  "Anshan",
			"email": "lahiru@gmail.com",
		}

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Put("http://localhost:8081/api/v1/students")

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 200, resp.StatusCode())
		assert.Equal(t, "{\"message\":\"Successfully updated\"}", resp.String())
	})

}

func TestIntegrationDeleteStudent(t *testing.T) {
	router, module, teardown := startup.TestServer()
	go router.Start(module.GetInstance().Env.ServerHost, module.GetInstance().Env.ServerPort)
	time.Sleep(1 * time.Second)

	defer teardown()

	// create a tempory student correctly
	createTemporaryStudent(t)

	t.Run("It should return a successfully deleted message", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		param := "lahiru@gmail.com"

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(param).
			Delete(fmt.Sprintf("http://localhost:8081/api/v1/students/%s", param))

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 200, resp.StatusCode())
		assert.Equal(t, "{\"message\":\"Successfully deleted\"}", resp.String())

	})

	t.Run("It should return an error message as student is not found", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		param := "lahiru@gmail.com"

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(param).
			Delete(fmt.Sprintf("http://localhost:8081/api/v1/students/%s", param))

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 404, resp.StatusCode())
		assert.Equal(t, "{\"error\":\"Student is not found\"}", resp.String())

	})

}

func TestIntegrationEnrollStudentToClass(t *testing.T) {
	router, module, teardown := startup.TestServer()
	go router.Start(module.GetInstance().Env.ServerHost, module.GetInstance().Env.ServerPort)
	time.Sleep(1 * time.Second)

	defer teardown()

	// create a tempory student correctly
	createTemporaryStudent(t)

	// create a tempory class correctly
	createTemporaryClass(t)

	t.Run("It should return a message contains successfully enrolled", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		body := map[string]interface{}{
			"student_email": "lahiru@gmail.com",
			"class_name":    "english-grade-12",
		}

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Put("http://localhost:8081/api/v1/students/enroll")

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 200, resp.StatusCode())
		assert.Equal(t, "{\"message\":\"Successfully enrolled\"}", resp.String())
	})
}

func TestIntegrationUnenrollStudentFromClass(t *testing.T) {
	router, module, teardown := startup.TestServer()
	go router.Start(module.GetInstance().Env.ServerHost, module.GetInstance().Env.ServerPort)
	time.Sleep(1 * time.Second)

	defer teardown()

	// create a tempory student correctly
	createTemporaryStudent(t)

	// create a tempory class correctly
	createTemporaryClass(t)

	// enroll student to a class
	enrollTemporaryStudentToTemporaryClass(t)

	t.Run("It should return a message contains successfully unenrolled", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		body := map[string]interface{}{
			"student_email": "lahiru@gmail.com",
			"class_name":    "english-grade-12",
		}

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Put("http://localhost:8081/api/v1/students/unenroll")

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 200, resp.StatusCode())
		assert.Equal(t, "{\"message\":\"Successfully unenrolled\"}", resp.String())
	})

}
