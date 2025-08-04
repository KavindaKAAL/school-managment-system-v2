package controller

import (
	"github.com/KavindaKAAL/school-management-system-v2/internal/core/port"
)

type baseController struct {
	basePath string
}

func NewBaseController(basePath string) port.BaseController {
	return &baseController{
		basePath: basePath,
	}
}

func (c *baseController) Path() string {
	return c.basePath
}
