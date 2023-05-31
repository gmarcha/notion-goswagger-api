package issue

import (
	"context"

	"github.com/gmarcha/notion-goswagger-api/internal/v1/goswagger/models"
	notion "github.com/jomei/notionapi"
)

type Service struct {
	notion *notion.Client
}

func (c *Service) FromJiraIDToNotionID(ctx context.Context, jiraIssueID string) (string, error) {

	body := &notion.DatabaseQueryRequest{
		Filter: &notion.PropertyFilter{
			Property: "Jira Issue ID",
			RichText: &notion.TextFilterCondition{
				Equals: jiraIssueID,
			},
		},
	}
	resp, err := c.notion.Database.Query(ctx, issuesDbID, body)
	if err != nil {
		return "", err
	}
	return string(resp.Results[0].ID), nil
}

func (c *Service) Create(ctx context.Context, issue *models.Issue) (*models.Issue, error) {

	body := &notion.PageCreateRequest{
		Parent: notion.Parent{DatabaseID: issuesDbID},
		Properties: notion.Properties{
			"Name":          notion.TitleProperty{Title: []notion.RichText{{Text: &notion.Text{Content: issue.Name}}}},
			"Status":        notion.SelectProperty{Select: notion.Option{Name: issue.Status}},
			"Type":          notion.MultiSelectProperty{MultiSelect: []notion.Option{{Name: issue.Type}}},
			"Jira Issue ID": notion.RichTextProperty{RichText: []notion.RichText{{Text: &notion.Text{Content: issue.JiraIssueID}}}},
		},
	}
	_, err := c.notion.Page.Create(ctx, body)
	if err != nil {
		return nil, err
	}
	return issue, nil
}

func (c *Service) Update(ctx context.Context, issueID string, issue *models.Issue) (*models.Issue, error) {

	pageID, err := c.FromJiraIDToNotionID(ctx, issueID)
	if err != nil {
		return nil, err
	}
	body := &notion.PageUpdateRequest{
		Properties: notion.Properties{
			"Name":   notion.TitleProperty{Title: []notion.RichText{{Text: &notion.Text{Content: issue.Name}}}},
			"Status": notion.SelectProperty{Select: notion.Option{Name: issue.Status}},
			"Type":   notion.MultiSelectProperty{MultiSelect: []notion.Option{{Name: issue.Type}}},
		},
	}
	_, err = c.notion.Page.Update(ctx, notion.PageID(pageID), body)
	if err != nil {
		return nil, err
	}
	return issue, nil
}

func (c *Service) Delete(ctx context.Context, issueID string) error {

	pageID, err := c.FromJiraIDToNotionID(ctx, issueID)
	if err != nil {
		return err
	}
	body := &notion.PageUpdateRequest{
		Archived:   true,
		Properties: notion.Properties{},
	}
	_, err = c.notion.Page.Update(ctx, notion.PageID(pageID), body)
	if err != nil {
		return err
	}
	return nil
}
