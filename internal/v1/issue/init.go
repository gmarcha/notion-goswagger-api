package issue

import (
	"os"

	"github.com/gmarcha/notion-goswagger-api/internal/v1/goswagger/restapi/operations"
	"github.com/gmarcha/notion-goswagger-api/internal/v1/log"
	notion "github.com/jomei/notionapi"
)

var (
	issuesDbID notion.DatabaseID
)

func init() {
	issuesDbID = notion.DatabaseID(os.Getenv("NOTION_ISSUES_DB_ID"))
	if issuesDbID == "" {
		log.Logger.Warn("Environment variable `NOTION_ISSUES_DB_ID` is empty")
	}
}

func Setup(api *operations.NotionAPI, client *notion.Client) {

	issueService := &Service{notion: client}

	api.IssueCreateHandler = &issueCreate{issue: issueService}
	api.IssueUpdateHandler = &issueUpdate{issue: issueService}
	api.IssueDeleteHandler = &issueDelete{issue: issueService}
}
