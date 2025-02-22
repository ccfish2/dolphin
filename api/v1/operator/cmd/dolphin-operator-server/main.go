// Code generated by go-swagger; DO NOT EDIT.

package main

import (
	"log"
	"os"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

	swaggerserversrv "github.com/ccfish2/dolphin/api/v1/operator/server"
	"github.com/ccfish2/dolphin/api/v1/operator/server/restapi"
)

// This file was generated by the swagger tool.
// Make sure not to overwrite this file after you generated it because all your edits would be lost!

func main() {

	swaggerSpec, err := loads.Embedded(swaggerserversrv.SwaggerJSON, swaggerserversrv.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	

	api := restapi.NewDolphinOperatorAPI(swaggerSpec)
	server := swaggerserversrv.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "dolphin operator"
	parser.LongDescription = "The API for the dolphin operator project"
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
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
		log.Fatalln(err)
	}

}
