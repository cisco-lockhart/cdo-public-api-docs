package services

import (
	"fmt"
	"github.com/cisco-lockhart/fcm-api-docs-generator/models"
	"github.com/pterm/pterm"
)

func MergeOpenApiSpecs(serviceSpecs map[string]*models.OpenAPI, config *models.Config) *models.OpenAPI {
	mergedSpec := models.OpenAPI{}
	mergedSpec.OpenAPI = "3.0.1"
	mergedSpec.Info = config.Info
	mergedSpec.Servers = config.Servers
	mergedSpec.Paths = make(map[string]interface{})
	mergedSpec.Components = models.Components{}
	mergedSpec.Components.Schemas = make(map[string]interface{})
	mergedSpec.Components.Responses = make(map[string]interface{})
	mergedSpec.Components.SecuritySchemes = config.SecuritySchemes

	for serviceName, spec := range serviceSpecs {
		spinner, _ := pterm.DefaultSpinner.Start(fmt.Sprintf("Merging OpenAPI spec for service: %s...", serviceName))
		// Merge paths
		for path, pathItem := range spec.Paths {
			if _, exists := mergedSpec.Paths[path]; exists {
				panic(fmt.Sprintf("duplicate path found: %s", path))
			}
			mergedSpec.Paths[path] = pathItem
		}

		// Merge components
		// Merge schemas
		for schemaName, schema := range spec.Components.Schemas {
			mergedSpec.Components.Schemas[schemaName] = schema
		}

		// Merge responses
		for responseName, response := range spec.Components.Responses {
			mergedSpec.Components.Responses[responseName] = response
		}

		spinner.Success(fmt.Sprintf("Merged OpenAPI spec for service: %s", serviceName))
	}

	return &mergedSpec
}
