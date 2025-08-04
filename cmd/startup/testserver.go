package startup

import (
	"net/http/httptest"

	"github.com/KavindaKAAL/school-management-system-v2/config"

	"github.com/KavindaKAAL/school-management-system-v2/internal/core/port"
)

type Teardown = func()

func TestServer() (port.Router, Module, Teardown) {
	env := config.NewEnv("../.test.env", false)
	router, module, shutdown := create(env)
	ts := httptest.NewServer(router.GetEngine())
	teardown := func() {
		ts.Close()
		shutdown()
	}
	return router, module, teardown
}
