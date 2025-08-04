package tests

import (
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
			Post("http://localhost:8081/api/v1/students")

		if err2 != nil {
			t.Fatalf("Request failed: %v", err)
		}

		assert.Equal(t, 409, resp2.StatusCode())
		assert.Equal(t, "{\"error\":\"Email is already used\"}", resp2.String())

	})
}

// func TestIntegrationGetStudentByEmail(t *testing.T) {
// 	ts, router, module, teardown := startup.TestServer()
// 	defer teardown()

// 	apikey, err := module.GetInstance().AuthService.CreateApiKey("test_key", 1, []model.Permission{"test"}, []string{"comment"})
// 	if err != nil {
// 		t.Fatalf("could not create apikey: %v", err)
// 	}

// 	body := `{"email":"test@abc.com","password":"123456","name":"test name"}`

// 	req, err := http.NewRequest("POST", "/auth/signup/basic", bytes.NewBuffer([]byte(body)))
// 	if err != nil {
// 		t.Fatalf("could not create request: %v", err)
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set(network.ApiKeyHeader, apikey.Key)

// 	rr := httptest.NewRecorder()
// 	router.GetEngine().ServeHTTP(rr, req)

// 	assert.Equal(t, http.StatusOK, rr.Code)
// 	assert.Contains(t, rr.Body.String(), `"message":"success"`)
// 	assert.Contains(t, rr.Body.String(), `"data"`)
// 	assert.Contains(t, rr.Body.String(), `"user"`)
// 	assert.Contains(t, rr.Body.String(), `"roles"`)
// 	assert.Contains(t, rr.Body.String(), `"tokens"`)

// 	_, err = module.GetInstance().AuthService.DeleteApiKey(apikey)
// 	if err != nil {
// 		t.Fatalf("could not delete apikey: %v", err)
// 	}

// 	_, err = module.GetInstance().UserService.DeleteUserByEmail("test@abc.com")
// 	if err != nil {
// 		t.Fatalf("could not delete user: %v", err)
// 	}
// }
