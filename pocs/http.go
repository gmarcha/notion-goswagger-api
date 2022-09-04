package pocs

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"sync"

	"github.com/gmarcha/notion-goswagger-api/internal/v1/goswagger/models"
)

type Parent struct {
	DatabaseId string `json:"database_id"`
}

type Text struct {
	Content string `json:"content"`
}

type Title struct {
	Text Text `json:"text"`
}

type Select struct {
	Name string `json:"name"`
}

type Relation struct {
	Id string `json:"id"`
}

type Sort struct {
	Property  string `json:"property"`
	Direction string `json:"direction"`
}

type ReqQueryDbBody struct {
	Sorts []Sort `json:"sorts"`
}

type TaskProperty struct {
	Title []Title `json:"title"`
}

type PriorityProperty struct {
	Select Select `json:"select"`
}

type IndexProperty struct {
	Number int `json:"number"`
}

type CampusProperty struct {
	Relations []Relation `json:"relation"`
}

type EpicsProperty struct {
	Relations []Relation `json:"relation"`
}

type CampusProperties struct {
	Task     TaskProperty     `json:"Task"`
	Priority PriorityProperty `json:"Priority"`
	Index    IndexProperty    `json:"Index"`
	Campus   CampusProperty   `json:"Campus"`
	Epics    EpicsProperty    `json:"Epics"`
}

type ReqCreatePageBody struct {
	Parent     Parent           `json:"parent"`
	Properties CampusProperties `json:"properties"`
}

type RespPage struct {
	Id string
}

type RespPageList struct {
	Object  string
	Results []RespPage
}

type RespPropertyItem struct {
	Title Title
}

type RespPropertyItemList struct {
	Object  string
	Results []RespPropertyItem
}

type Service struct {
	http *http.Client
}

var (
	tasksDbId           string
	onboardingTasksDbId string
)

func init() {
	tasksDbId = os.Getenv("NOTION_TASKS_DB_ID")
	onboardingTasksDbId = os.Getenv("NOTION_ONBOARDING_TASKS_DB_ID")
}

func (c *Service) CreateTasks(ctx context.Context, campusId string) ([]*models.Task, error) {

	sorts := &ReqQueryDbBody{Sorts: []Sort{{Property: "Step", Direction: "ascending"}}}
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(sorts)
	if err != nil {
		return nil, err
	}

	url := "https://api.notion.com/v1/databases/" + onboardingTasksDbId + "/query"
	res, err := c.http.Post(url, "application/json", b)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	list := &RespPageList{}
	err = json.NewDecoder(res.Body).Decode(list)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup

	for i, page := range list.Results {

		wg.Add(1)
		go func(index int, page RespPage) {

			defer wg.Done()

			url := "https://api.notion.com/v1/pages/" + page.Id + "/properties/title"
			res, err := c.http.Get(url)
			if err != nil {
				return
			}
			defer res.Body.Close()

			list := &RespPropertyItemList{}
			err = json.NewDecoder(res.Body).Decode(list)
			if err != nil {
				return
			}

			task := list.Results[0].Title.Text.Content

			reqCreatePageBody := &ReqCreatePageBody{
				Parent: Parent{DatabaseId: tasksDbId},
				Properties: CampusProperties{
					Task:     TaskProperty{Title: []Title{{Text: Text{Content: task}}}},
					Priority: PriorityProperty{Select: Select{Name: "P1"}},
					Index:    IndexProperty{Number: index + 1},
					Campus:   CampusProperty{Relations: []Relation{{Id: campusId}}},
					Epics:    EpicsProperty{Relations: []Relation{{Id: campusId}}},
				},
			}
			b := new(bytes.Buffer)
			err = json.NewEncoder(b).Encode(reqCreatePageBody)
			if err != nil {
				return
			}

			url = "https://api.notion.com/v1/pages"
			res, err = c.http.Post(url, "application/json", b)
			if err != nil {
				return
			}
			defer res.Body.Close()
		}(i, page)
	}
	wg.Wait()

	return []*models.Task{}, nil
}
