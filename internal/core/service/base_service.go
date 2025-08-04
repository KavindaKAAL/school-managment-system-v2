package service

import (
	"context"

	"github.com/KavindaKAAL/school-management-system-v2/internal/core/port"
)

type baseService struct {
	context context.Context
}

func NewBaseService() port.BaseService {
	return &baseService{
		context: context.Background(),
	}
}

func (s *baseService) Context() context.Context {
	return s.context
}
