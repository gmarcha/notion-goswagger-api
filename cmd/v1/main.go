package main

import (
	"os"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

	"github.com/gmarcha/notion-goswagger-api/internal/v1/goswagger/restapi"
	"github.com/gmarcha/notion-goswagger-api/internal/v1/goswagger/restapi/operations"
	"github.com/gmarcha/notion-goswagger-api/internal/v1/log"
)

func main() {

	defer log.Logger.Sync()

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Logger.Fatal(err.Error())
	}

	api := operations.NewNotionAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Notion API"
	parser.LongDescription = "An API to automate content creation on Notion."
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Logger.Fatal(err.Error())
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Logger.Fatal(err.Error())
	}

}
