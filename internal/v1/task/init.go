package task

import (
	"github.com/gmarcha/notion-goswagger-api/internal/v1/goswagger/restapi/operations"
	"github.com/jomei/notionapi"
)

func Setup(api *operations.NotionAPI, client *notionapi.Client) {

	taskService := &Service{notion: client}

	api.TaskCampusCreateHandler = &campusCreate{task: taskService}
	api.TaskPoolCreateHandler = &poolCreate{task: taskService}
}
