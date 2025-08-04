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

var createTemporaryTeacher = func(t *testing.T) {

	client := resty.New()
	client.SetRetryCount(5)
	client.SetRetryWaitTime(2 * time.Second)

	body := map[string]interface{}{
		"name":  "alice",
		"email": "alice@gmail.com",
	}

	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post("http://localhost:8081/api/v1/teachers")

	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
}

var assignTemporaryTeacherToTemporaryClass = func(t *testing.T) {

	client := resty.New()
	client.SetRetryCount(5)
	client.SetRetryWaitTime(2 * time.Second)

	body := map[string]interface{}{
		"teacher_email": "alice@gmail.com",
		"class_name":    "english-grade-12",
	}

	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Put("http://localhost:8081/api/v1/classes/assign")

	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
}

func TestIntegrationCreateClass(t *testing.T) {
	router, module, teardown := startup.TestServer()
	go router.Start(module.GetInstance().Env.ServerHost, module.GetInstance().Env.ServerPort)
	time.Sleep(1 * time.Second)

	defer teardown()

	t.Run("It should return successfully created message when creating a new class successfully", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		body := map[string]interface{}{
			"name":    "english-grade-12",
			"subject": "English",
		}

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Post("http://localhost:8081/api/v1/classes")

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 201, resp.StatusCode())
		assert.Equal(t, "{\"message\":\"Successfully created\"}", resp.String())

	})

	t.Run("It should return error message when creating a new class using already existing name", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		// 1st rest call
		createTemporaryClass(t)

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		body := map[string]interface{}{
			"name":    "english-grade-12",
			"subject": "English-12",
		}
		// 2nd rest call
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Post("http://localhost:8081/api/v1/classes")

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 409, resp.StatusCode())
		assert.Equal(t, "{\"error\":\"Class name is already used\"}", resp.String())

	})

	t.Run("It should return error message when creating a new class using missing required field class subject", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		body := map[string]interface{}{
			"name": "english-grade-12",
			"age":  25,
		}

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Post("http://localhost:8081/api/v1/classes")

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 400, resp.StatusCode())
		assert.Equal(t, "{\"error\":\"Invalid request body\"}", resp.String())

	})

	t.Run("It should return error message when creating a new class using missing required field class name", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		body := map[string]interface{}{
			"subject": "English-12",
			"age":     25,
		}

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Post("http://localhost:8081/api/v1/classes")

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 400, resp.StatusCode())
		assert.Equal(t, "{\"error\":\"Invalid request body\"}", resp.String())

	})

	t.Run("It should return error message when creating a new class using incorrect value for the name field", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		body := map[string]interface{}{
			"name":    12,
			"subject": "English-12",
		}

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Post("http://localhost:8081/api/v1/classes")

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 400, resp.StatusCode())
		assert.Equal(t, "{\"error\":\"Invalid type for field name\"}", resp.String())

	})

	t.Run("It should return error message when creating a new class using incorrect value for the subject field", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		body := map[string]interface{}{
			"name":    "english-grade-12",
			"subject": 12,
		}

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Post("http://localhost:8081/api/v1/classes")

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 400, resp.StatusCode())
		assert.Equal(t, "{\"error\":\"Invalid type for field subject\"}", resp.String())

	})

	t.Run("It should return error message when creating a new class using incorrect request method", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		body := map[string]interface{}{
			"name":    "english-grade-12",
			"subject": "English-12",
		}

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Patch("http://localhost:8081/api/v1/classes")

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 404, resp.StatusCode())
		assert.Equal(t, "{\"error\":\"Route not found\"}", resp.String())

	})

	t.Run("It should return error message when creating a new class using incorrect json format", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		invalidBody := `{
			"name":    "english-grade-12",
			"subject": "English-12",
		}`

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(invalidBody).
			Post("http://localhost:8081/api/v1/classes")

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 400, resp.StatusCode())
		assert.Equal(t, "{\"error\":\"Malformed JSON: Syntax error\"}", resp.String())

	})

}

func TestIntegrationGetSpecificClass(t *testing.T) {
	router, module, teardown := startup.TestServer()
	go router.Start(module.GetInstance().Env.ServerHost, module.GetInstance().Env.ServerPort)
	time.Sleep(1 * time.Second)

	defer teardown()

	// create a tempory class correctly
	createTemporaryClass(t)

	t.Run("It should return a message contains the class details", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		param := "english-grade-12"

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(param).
			Get(fmt.Sprintf("http://localhost:8081/api/v1/classes/%s", param))

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 200, resp.StatusCode())

	})

	t.Run("It should return an error message of class not found", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		param := "english-grade-12"

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(param).
			Get(fmt.Sprintf("http://localhost:8081/api/v1/classes/%s", param))

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 404, resp.StatusCode())
		assert.Equal(t, "{\"error\":\"Class is not found\"}", resp.String())

	})
}

func TestIntegrationUpdateClass(t *testing.T) {
	router, module, teardown := startup.TestServer()
	go router.Start(module.GetInstance().Env.ServerHost, module.GetInstance().Env.ServerPort)
	time.Sleep(1 * time.Second)

	defer teardown()

	// create a tempory class correctly
	createTemporaryClass(t)

	t.Run("It should return a message contains successfully updated", func(t *testing.T) {
		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		body := map[string]interface{}{
			"name":    "english-grade-12",
			"subject": "Maths",
		}

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Put("http://localhost:8081/api/v1/classes")

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 200, resp.StatusCode())
		assert.Equal(t, "{\"message\":\"Successfully updated\"}", resp.String())
	})

}

func TestIntegrationDeleteClass(t *testing.T) {
	router, module, teardown := startup.TestServer()
	go router.Start(module.GetInstance().Env.ServerHost, module.GetInstance().Env.ServerPort)
	time.Sleep(1 * time.Second)

	defer teardown()

	// create a tempory class correctly
	createTemporaryClass(t)

	t.Run("It should return a successfully deleted message", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		param := "english-grade-12"

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(param).
			Delete(fmt.Sprintf("http://localhost:8081/api/v1/classes/%s", param))

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 200, resp.StatusCode())
		assert.Equal(t, "{\"message\":\"Successfully deleted\"}", resp.String())

	})

	t.Run("It should return an error message as class is not found", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		param := "english-grade-12"

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(param).
			Delete(fmt.Sprintf("http://localhost:8081/api/v1/classes/%s", param))

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 404, resp.StatusCode())
		assert.Equal(t, "{\"error\":\"Class is not found\"}", resp.String())

	})

}

func TestIntegrationAssignTeacherToClass(t *testing.T) {
	router, module, teardown := startup.TestServer()
	go router.Start(module.GetInstance().Env.ServerHost, module.GetInstance().Env.ServerPort)
	time.Sleep(1 * time.Second)

	defer teardown()

	// create a tempory teacher correctly
	createTemporaryTeacher(t)

	// create a tempory class correctly
	createTemporaryClass(t)

	t.Run("It should return a message contains successfully assigned", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		body := map[string]interface{}{
			"teacher_email": "alice@gmail.com",
			"class_name":    "english-grade-12",
		}

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Put("http://localhost:8081/api/v1/classes/assign")

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 200, resp.StatusCode())
		assert.Equal(t, "{\"message\":\"Successfully assigned\"}", resp.String())
	})

}

func TestIntegrationUnassignTeacherFromClass(t *testing.T) {
	router, module, teardown := startup.TestServer()
	go router.Start(module.GetInstance().Env.ServerHost, module.GetInstance().Env.ServerPort)
	time.Sleep(1 * time.Second)

	defer teardown()

	// create a tempory teacher correctly
	createTemporaryTeacher(t)

	// create a tempory class correctly
	createTemporaryClass(t)

	// assign teacher to a class
	assignTemporaryTeacherToTemporaryClass(t)

	t.Run("It should return a message contains successfully unassigned", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		body := map[string]interface{}{
			"teacher_email": "alice@gmail.com",
			"class_name":    "english-grade-12",
		}

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Put("http://localhost:8081/api/v1/classes/unassign")

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 200, resp.StatusCode())
		assert.Equal(t, "{\"message\":\"Successfully unassigned\"}", resp.String())
	})

}
