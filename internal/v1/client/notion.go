package client

import (
	"github.com/gmarcha/notion-goswagger-api/internal/v1/goswagger/restapi/operations"
	"github.com/jomei/notionapi"
)

func Notion(api *operations.NotionAPI) *notionapi.Client {
	return notionapi.NewClient(notionapi.Token(notionToken))
}
