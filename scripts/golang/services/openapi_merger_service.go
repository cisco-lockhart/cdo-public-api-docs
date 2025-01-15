package services

import (
	"errors"
	"fmt"
	"github.com/cisco-lockhart/cloud-fw-mgr-api-docs/models"
	"github.com/pterm/pterm"
	"strings"
)

func MergeOpenApiSpecs(serviceSpecs map[string]*models.OpenAPI, config *models.Config) (*models.OpenAPI, error) {
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
		var prefixToAdd *string
		for _, service := range config.Services {
			if service.Name == serviceName && service.PrefixToAdd != nil {
				trimmedPrefix := strings.TrimSuffix(*service.PrefixToAdd, "/")
				if !strings.HasPrefix(trimmedPrefix, "/") {
					trimmedPrefix = "/" + trimmedPrefix
				}
				prefixToAdd = &trimmedPrefix
				break
			}
		}

		spinner, _ := pterm.DefaultSpinner.Start(fmt.Sprintf("Merging OpenAPI spec for service: %s...", serviceName))
		// Merge paths
		for path, pathItem := range spec.Paths {
			if prefixToAdd != nil {
				path = *prefixToAdd + path
			}
			if _, exists := mergedSpec.Paths[path]; exists {
				return nil, errors.New(fmt.Sprintf("duplicate path found: %s", path))
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

	return &mergedSpec, nil
}
