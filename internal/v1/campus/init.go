package campus

import (
	"os"

	"github.com/gmarcha/notion-goswagger-api/internal/v1/goswagger/restapi/operations"
	"github.com/gmarcha/notion-goswagger-api/internal/v1/log"
	notion "github.com/jomei/notionapi"
)

var (
	campusDbID notion.DatabaseID
	dbToSync   notion.DatabaseID
)

func init() {
	campusDbID = notion.DatabaseID(os.Getenv("NOTION_CAMPUS_DB_ID"))
	if campusDbID == "" {
		log.Logger.Warn("Environment variable `NOTION_CAMPUS_DB_ID` is empty")
	}
	dbToSync = notion.DatabaseID("26619d6d928144f59b9d5465c1273163")
}

func Setup(api *operations.NotionAPI, client *notion.Client) {

	campusService := &Service{notion: client}

	api.CampusSyncHandler = &sync{campus: campusService}
}
