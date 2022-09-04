package task

import (
	"context"

	"github.com/gmarcha/notion-goswagger-api/internal/v1/goswagger/models"
	"github.com/gmarcha/notion-goswagger-api/internal/v1/goswagger/restapi/operations/task"
	"github.com/go-openapi/runtime/middleware"
)

type campusCreate struct {
	task *Service
}

func (s *campusCreate) Handle(params task.CampusCreateParams) middleware.Responder {
	ctx := context.Background()
	res, err := s.task.CampusCreate(ctx, params.ID)
	if err != nil {
		return task.NewCampusCreateInternalServerError().WithPayload(
			&models.Error{Code: 500, Message: err.Error()})
	}
	return task.NewCampusCreateCreated().WithPayload(res)
}

type poolCreate struct {
	task *Service
}

func (s *poolCreate) Handle(params task.PoolCreateParams) middleware.Responder {
	ctx := context.Background()
	res, err := s.task.PoolCreate(ctx, params.ID)
	if err != nil {
		return task.NewPoolCreateInternalServerError().WithPayload(
			&models.Error{Code: 500, Message: err.Error()})
	}
	return task.NewPoolCreateCreated().WithPayload(res)
}
