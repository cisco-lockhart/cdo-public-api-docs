/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/cisco-lockhart/cloud-fw-mgr-api-docs/models"
	"github.com/cisco-lockhart/cloud-fw-mgr-api-docs/services"
	"github.com/pterm/pterm"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// generateCdfmcApiDocsCmd represents the generateCdfmcApiDocs command
var generateCdfmcApiDocsCmd = &cobra.Command{
	Use:   "generate-cdfmc-api-docs",
	Short: "Generate cdFMC API docs to include into the documentation",
	Long: `Generate cdFMC API docs, rewriting links to reflect the CDO URLs.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		configUrl, err := cmd.Flags().GetString("config")
		if err != nil {
			pterm.Error.Println("Failed to get config URL", err)
			os.Exit(1)
		}
		config, err := loadConfig(configUrl)
		if err != nil {
			os.Exit(1)
		}
		outputFile, err := cmd.Flags().GetString("output")
		if err != nil {
			pterm.Error.Println("Failed to get output file path", err)
			os.Exit(1)
		}

		openApiSpec, err := loadOpenApiSpec(config.CdFmc.URL)
		if err != nil {
			os.Exit(1)
		}
		modifiedOpenApiSpec := modifySpec(openApiSpec, config)
		err = writeModifiedCdFmcSpecToFile(modifiedOpenApiSpec, outputFile)
		if err != nil {
			os.Exit(1)
		}
	},
}

func writeModifiedCdFmcSpecToFile(modifiedOpenApiSpec *models.OpenAPI, outputFile string) error {
	spinner, _ := pterm.DefaultSpinner.Start("Writing modified cdFMC OpenAPI spec to file...")
	file, err := os.Create(outputFile)
	if err != nil {
		spinner.Fail(fmt.Sprintf("failed to create file: %v\n", err))
		return err
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	err = encoder.Encode(modifiedOpenApiSpec)
	if err != nil {
		spinner.Fail(fmt.Sprintf("failed to write OpenAPI spec to file: %v\n", err))
		return err
	}
	spinner.Success(fmt.Sprintf("Modified cdFMC OpenAPI spec written to file: %s", outputFile))

	return nil
}

func modifySpec(openApiSpec *models.OpenAPI, config *models.Config) *models.OpenAPI {
	spinner, _ := pterm.DefaultSpinner.Start("Rewriting paths and modifying servers...")
	openApiSpec.OpenAPI = "3.0.1"
	openApiSpec.Servers = config.Servers
	openApiSpec.Info = config.CdFmcInfo
	openApiSpec.Components.SecuritySchemes = config.SecuritySchemes

	newPaths := make(map[string]interface{})
	for path, pathItem := range openApiSpec.Paths {
		newPath := "/v1/cdfmc" + path
		newPaths[newPath] = pathItem
	}
	openApiSpec.Paths = newPaths
	spinner.Success("Paths and servers modified successfully")
	return openApiSpec
}

func loadConfig(configUrl string) (*models.Config, error) {
	spinner, _ := pterm.DefaultSpinner.Start("Loading configuration file...")
	config, err := services.LoadConfig(configUrl)
	if err != nil {
		spinner.Fail(fmt.Sprintf("failed to load configuration file: %v\n", err))
		return nil, err
	}
	spinner.Success(fmt.Sprintf("Configuration file loaded successfully: %s", configUrl))

	return config, nil
}

func loadOpenApiSpec(url string) (*models.OpenAPI, error) {
	spinner, _ := pterm.DefaultSpinner.Start(fmt.Sprintf("Loading OpenAPI spec for cdFMC"))
	openApiSpec, err := services.LoadOpenApiJson(url)
	if err != nil {
		spinner.Fail(fmt.Sprintf("failed to load OpenAPI spec for cdFMC, error: %v\n", err))
		return nil, err
	}
	spinner.Success(fmt.Sprintf("OpenAPI spec loaded successfully for cdFMC"))

	return openApiSpec, err
}

func init() {
	rootCmd.AddCommand(generateCdfmcApiDocsCmd)

	currDir, err := os.Getwd()
	if err != nil {
		pterm.Error.Println("Failed to get current working directory", err)
	}
	defaultOutputFile := filepath.Join(currDir, "cdfmc-openapi.yaml")
	generateCdfmcApiDocsCmd.Flags().StringP("output", "o", defaultOutputFile, fmt.Sprintf("Specify the output file path to write the merged OpenAPI spec to (default: %s)", defaultOutputFile))
}
