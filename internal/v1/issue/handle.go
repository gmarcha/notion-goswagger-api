package issue

import (
	"context"

	"github.com/gmarcha/notion-goswagger-api/internal/v1/goswagger/models"
	"github.com/gmarcha/notion-goswagger-api/internal/v1/goswagger/restapi/operations/issue"
	"github.com/gmarcha/notion-goswagger-api/internal/v1/log"
	"github.com/go-openapi/runtime/middleware"
)

type issueCreate struct {
	issue *Service
}

func (s *issueCreate) Handle(params issue.CreateParams) middleware.Responder {
	ctx := context.Background()
	resp, err := s.issue.Create(ctx, params.Issue)
	if err != nil {
		log.Logger.Error(err.Error())
		return issue.NewCreateInternalServerError().WithPayload(
			&models.Error{Code: 500, Message: err.Error()})
	}
	log.Logger.Debug("Issue created")
	return issue.NewCreateCreated().WithPayload(resp)
}

type issueUpdate struct {
	issue *Service
}

func (s *issueUpdate) Handle(params issue.UpdateParams) middleware.Responder {
	ctx := context.Background()
	res, err := s.issue.Update(ctx, params.ID, params.Issue)
	if err != nil {
		log.Logger.Error(err.Error())
		return issue.NewUpdateInternalServerError().WithPayload(
			&models.Error{Code: 500, Message: err.Error()})
	}
	log.Logger.Debug("Issue updated")
	return issue.NewUpdateOK().WithPayload(res)
}

type issueDelete struct {
	issue *Service
}

func (s *issueDelete) Handle(params issue.DeleteParams) middleware.Responder {
	ctx := context.Background()
	err := s.issue.Delete(ctx, params.ID)
	if err != nil {
		log.Logger.Error(err.Error())
		return issue.NewDeleteInternalServerError().WithPayload(
			&models.Error{Code: 500, Message: err.Error()})
	}
	log.Logger.Debug("Issue deleted")
	return issue.NewDeleteNoContent()
}
