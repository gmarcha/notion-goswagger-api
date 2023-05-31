package campus

import (
	"context"

	notion "github.com/jomei/notionapi"
)

type PageMap = map[notion.ObjectID]notion.Page

type Service struct {
	notion *notion.Client
}

func (c *Service) FetchInMap(ctx context.Context, dbID notion.DatabaseID) (PageMap, error) {

	newList, err := c.notion.Database.Query(ctx, dbID, nil)
	if err != nil {
		return nil, err
	}

	newMap := map[notion.ObjectID]notion.Page{}
	for _, page := range newList.Results {
		newMap[page.ID] = page
	}
	return newMap, nil
}

func (c *Service) Sync(ctx context.Context, dbToSync notion.DatabaseID) error {

	body := &notion.DatabaseQueryRequest{
		// Filter: &notion.PropertyFilter{
		// 	Property: "Jira Issue ID",
		// 	RichText: &notion.TextFilterCondition{
		// 		Equals: jiraIssueID,
		// 	},
		// },
	}
	currentPageMap, err := c.FetchInMap(ctx, campusDbID)
	if err != nil {
		return err
	}
	oldPageMap, err := c.FetchInMap(ctx, dbToSync)
	if err != nil {
		return err
	}
	// for pageID, page := range currentPageMap {
	// 	if oldPageMap[pageID]
	// }
	return nil
}
