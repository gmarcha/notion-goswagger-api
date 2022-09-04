package client

import (
	"os"
)

var notionToken string

func init() {
	notionToken = os.Getenv("NOTION_TOKEN")
}
