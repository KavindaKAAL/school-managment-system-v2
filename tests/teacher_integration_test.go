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

func TestIntegrationCreatetTeacher(t *testing.T) {
	router, module, teardown := startup.TestServer()
	go router.Start(module.GetInstance().Env.ServerHost, module.GetInstance().Env.ServerPort)
	time.Sleep(1 * time.Second)

	defer teardown()

	t.Run("It should return successfully created message when creating a new teacher successfully", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		body := map[string]interface{}{
			"name":  "Alice",
			"email": "alice@gmail.com",
		}

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Post("http://localhost:8081/api/v1/teachers")

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 201, resp.StatusCode())
		assert.Equal(t, "{\"message\":\"Successfully created\"}", resp.String())

	})

	t.Run("It should return error message when creating a new teacher using already existing email", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		body := map[string]interface{}{
			"name":  "Alice",
			"email": "alice@gmail.com",
		}

		_, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Post("http://localhost:8081/api/v1/teachers")

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		body2 := map[string]interface{}{
			"name":  "Kavinda",
			"email": "alice@gmail.com",
		}
		// 2nd rest call
		resp2, err2 := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body2).
			Post("http://localhost:8081/api/v1/teachers")

		if err2 != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 409, resp2.StatusCode())
		assert.Equal(t, "{\"error\":\"Email is already used\"}", resp2.String())

	})
}

func TestIntegrationGetSpecificTeacher(t *testing.T) {
	router, module, teardown := startup.TestServer()
	go router.Start(module.GetInstance().Env.ServerHost, module.GetInstance().Env.ServerPort)
	time.Sleep(1 * time.Second)

	defer teardown()

	// create a tempory teacher correctly
	createTemporaryTeacher(t)

	t.Run("It should return a message contains the teacher details", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		param := "alice@gmail.com"

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(param).
			Get(fmt.Sprintf("http://localhost:8081/api/v1/teachers/%s", param))

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 200, resp.StatusCode())

	})

	t.Run("It should return an error message of teacher not found", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		param := "alice@gmail.com"

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(param).
			Get(fmt.Sprintf("http://localhost:8081/api/v1/teachers/%s", param))

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 404, resp.StatusCode())
		assert.Equal(t, "{\"error\":\"Teacher is not found\"}", resp.String())

	})
}

func TestIntegrationUpdateTeacher(t *testing.T) {
	router, module, teardown := startup.TestServer()
	go router.Start(module.GetInstance().Env.ServerHost, module.GetInstance().Env.ServerPort)
	time.Sleep(1 * time.Second)

	defer teardown()

	// create a tempory teacher correctly
	createTemporaryTeacher(t)

	t.Run("It should return a message contains successfully updated", func(t *testing.T) {
		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		body := map[string]interface{}{
			"name":  "Anshan",
			"email": "alice@gmail.com",
		}

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Put("http://localhost:8081/api/v1/teachers")

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 200, resp.StatusCode())
		assert.Equal(t, "{\"message\":\"Successfully updated\"}", resp.String())
	})
}

func TestIntegrationDeleteTeacher(t *testing.T) {
	router, module, teardown := startup.TestServer()
	go router.Start(module.GetInstance().Env.ServerHost, module.GetInstance().Env.ServerPort)
	time.Sleep(1 * time.Second)

	defer teardown()

	// create a tempory teacher correctly
	createTemporaryTeacher(t)

	t.Run("It should return a successfully deleted message", func(t *testing.T) {
		defer utils.ClearDatabase(module.GetInstance().DB.GetInstance())

		client := resty.New()
		client.SetRetryCount(5)
		client.SetRetryWaitTime(2 * time.Second)

		param := "alice@gmail.com"

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(param).
			Delete(fmt.Sprintf("http://localhost:8081/api/v1/teachers/%s", param))

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

		param := "alice@gmail.com"

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(param).
			Delete(fmt.Sprintf("http://localhost:8081/api/v1/teachers/%s", param))

		if err != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 404, resp.StatusCode())
		assert.Equal(t, "{\"error\":\"Teacher is not found\"}", resp.String())

	})

}
