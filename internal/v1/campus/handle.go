package campus

import (
	"context"

	"github.com/gmarcha/notion-goswagger-api/internal/v1/goswagger/models"
	"github.com/gmarcha/notion-goswagger-api/internal/v1/goswagger/restapi/operations/campus"
	"github.com/gmarcha/notion-goswagger-api/internal/v1/log"
	"github.com/go-openapi/runtime/middleware"
)

type sync struct {
	campus *Service
}

func (s *sync) Handle(params campus.SyncParams) middleware.Responder {
	ctx := context.Background()
	err := s.campus.Sync(ctx, dbToSync)
	if err != nil {
		log.Logger.Error(err.Error())
		return campus.NewSyncInternalServerError().WithPayload(
			&models.Error{Code: 500, Message: err.Error()})
	}
	log.Logger.Debug("Campuses synchronized")
	return campus.NewSyncNoContent()
}
