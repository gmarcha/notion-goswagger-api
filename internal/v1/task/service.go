package task

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/gmarcha/notion-goswagger-api/internal/v1/goswagger/models"
	notion "github.com/jomei/notionapi"
)

type Service struct {
	notion *notion.Client
}

var (
	tasksDbId           notion.DatabaseID
	onboardingTasksDbId notion.DatabaseID
	poolTasksDbId       notion.DatabaseID
)

func init() {
	tasksDbId = notion.DatabaseID(os.Getenv("NOTION_TASKS_DB_ID"))
	onboardingTasksDbId = notion.DatabaseID(os.Getenv("NOTION_ONBOARDING_TASKS_DB_ID"))
	poolTasksDbId = notion.DatabaseID(os.Getenv("NOTION_POOL_TASKS_DB_ID"))
}

func (c *Service) CampusCreate(ctx context.Context, campusId string) ([]*models.Task, error) {

	body := &notion.DatabaseQueryRequest{
		Sorts: []notion.SortObject{{
			Property:  "Step",
			Direction: "ascending",
		}},
	}
	resp, err := c.notion.Database.Query(ctx, onboardingTasksDbId, body)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup

	for i, page := range resp.Results {

		wg.Add(1)
		go func(index int, page notion.Page) {

			defer wg.Done()

			resp, err := c.notion.Page.Get(ctx, notion.PageID(page.ID))
			if err != nil {
				return
			}

			task := resp.Properties["Task"].(*notion.TitleProperty)
			epic := resp.Properties["Epic"].(*notion.RelationProperty)
			priority := resp.Properties["Priority"].(*notion.SelectProperty).Select.Name
			step := resp.Properties["Step"].(*notion.NumberProperty)

			body := &notion.PageCreateRequest{
				Parent: notion.Parent{
					DatabaseID: tasksDbId,
				},
				Properties: map[string]notion.Property{
					"Task":     task,
					"Campus":   notion.RelationProperty{Relation: []notion.Relation{{ID: notion.PageID(campusId)}}},
					"Epics":    epic,
					"Priority": notion.SelectProperty{Select: notion.Option{Name: priority}},
					"Index":    step,
				},
			}
			_, err = c.notion.Page.Create(ctx, body)
			if err != nil {
				fmt.Println(err)
				return
			}
		}(i, page)
	}
	wg.Wait()

	return []*models.Task{}, nil
}

func (c *Service) PoolCreate(ctx context.Context, eventId string) ([]*models.Task, error) {

	body := &notion.DatabaseQueryRequest{
		Sorts: []notion.SortObject{{
			Property:  "Step",
			Direction: "ascending",
		}},
	}
	resp, err := c.notion.Database.Query(ctx, poolTasksDbId, body)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup

	for i, page := range resp.Results {

		wg.Add(1)
		go func(index int, page notion.Page) {

			defer wg.Done()

			resp, err := c.notion.Page.Get(ctx, notion.PageID(page.ID))
			if err != nil {
				return
			}

			task := resp.Properties["Task"].(*notion.TitleProperty)
			epic := resp.Properties["Epic"].(*notion.RelationProperty)
			priority := resp.Properties["Priority"].(*notion.SelectProperty).Select.Name
			step := resp.Properties["Step"].(*notion.NumberProperty)

			body := &notion.PageCreateRequest{
				Parent: notion.Parent{
					DatabaseID: tasksDbId,
				},
				Properties: map[string]notion.Property{
					"Task":     task,
					"Event":    notion.RelationProperty{Relation: []notion.Relation{{ID: notion.PageID(eventId)}}},
					"Epics":    epic,
					"Priority": notion.SelectProperty{Select: notion.Option{Name: priority}},
					"Index":    step,
				},
			}
			_, err = c.notion.Page.Create(ctx, body)
			if err != nil {
				fmt.Println(err)
				return
			}
		}(i, page)
	}
	wg.Wait()

	return []*models.Task{}, nil
}
