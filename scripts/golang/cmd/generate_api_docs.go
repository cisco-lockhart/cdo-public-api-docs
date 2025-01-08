/*
Copyright © 2025 Cisco Systems
*/
package cmd

import (
	"fmt"
	"github.com/cisco-lockhart/fcm-api-docs-generator/models"
	"github.com/cisco-lockhart/fcm-api-docs-generator/services"
	"github.com/pterm/pterm"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate-api-docs",
	Short: "Combine the API documentation from each microservice that constitutes the FCM API into one single API spec.",
	Long: `This command downloads OpenAPI specs for all of the microservices that constitute 
			the Firewall public API and merges them together (including by removing schemas shared 
			between the microservices).`,
	Run: func(cmd *cobra.Command, args []string) {
		configUrl, err := cmd.Flags().GetString("config")
		if err != nil {
			pterm.Error.Println("Failed to get config URL", err)
			os.Exit(1)
		}
		outputFile, err := cmd.Flags().GetString("output")
		if err != nil {
			pterm.Error.Println("Failed to get output file path", err)
			os.Exit(1)
		}

		spinner, _ := pterm.DefaultSpinner.Start("Loading configuration file...")
		config := services.LoadConfig(configUrl)
		spinner.Success(fmt.Sprintf("Configuration file loaded successfully: %s", configUrl))

		serviceSpecs := make(map[string]*models.OpenAPI)
		for _, service := range config.Services {
			spinner, _ = pterm.DefaultSpinner.Start(fmt.Sprintf("Loading OpenAPI spec for service: %s...", service.Name))
			openApiSpec, err := services.LoadOpenApi(service.URL)
			if err != nil {
				spinner.Fail(fmt.Sprintf("failed to load OpenAPI spec for service: %s, error: %v\n", service.Name, err))
				os.Exit(1)
			}
			serviceSpecs[service.Name] = openApiSpec
			spinner.Success(fmt.Sprintf("OpenAPI spec loaded successfully for service: %s", service.Name))
		}

		spinner, _ = pterm.DefaultSpinner.Start("Merging OpenAPI specs...")
		mergedSpec, err := services.MergeOpenApiSpecs(serviceSpecs, config)
		if err != nil {
			spinner.Fail(fmt.Sprintf("failed to merge OpenAPI specs: %v\n", err))
			os.Exit(1)
		}
		spinner.Success("OpenAPI specs merged successfully")

		spinner, _ = pterm.DefaultSpinner.Start(fmt.Sprintf("Marshalling merged spec to YAML and writing to %s...", outputFile))
		yamlData, err := yaml.Marshal(mergedSpec)
		if err != nil {
			spinner.Fail(fmt.Sprintf("failed to marshal merged spec: %v\n", err))
		}
		// Write the YAML data to a file
		err = os.WriteFile(outputFile, yamlData, 0644)
		if err != nil {
			spinner.Fail(fmt.Sprintf("failed to write merged spec to file: %v\n", err))
			os.Exit(1)
		}
		spinner.Success("Merged spec marshalled to YAML successfully")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	currDir, err := os.Getwd()
	if err != nil {
		pterm.Error.Println("Failed to get current working directory", err)
	}
	defaultOutputFile := filepath.Join(currDir, "openapi.yaml")
	generateCmd.Flags().StringP("output", "o", defaultOutputFile, fmt.Sprintf("Specify the output file path to write the merged OpenAPI spec to (default: %s)", defaultOutputFile))
	generateCmd.Flags().BoolP("generate-sdks", "s", true, "Generate SDKs from the merged OpenAPI spec")
	generateCmd.Flags().BoolP("generate-postman-collection", "d", true, "Generate documentation from the merged OpenAPI spec")
}
