package client

import (
	"net/http"

	"github.com/gmarcha/notion-goswagger-api/internal/v1/goswagger/restapi/operations"
)

type NotionTransport struct{}

func (t *NotionTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+notionToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Notion-Version", "2022-06-28")
	return http.DefaultTransport.RoundTrip(req)
}

func Http(api *operations.NotionAPI) *http.Client {
	return &http.Client{Transport: &NotionTransport{}}
}
