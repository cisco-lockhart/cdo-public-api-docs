/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/cisco-lockhart/cloud-fw-mgr-api-docs/models"
	"github.com/cisco-lockhart/cloud-fw-mgr-api-docs/services"
	"github.com/pterm/pterm"
	"gopkg.in/yaml.v3"
	"os"

	"github.com/spf13/cobra"
)

// generateSdksCmd represents the generateSdks command
var generateSdksCmd = &cobra.Command{
	Use:   "generate-sdks <path-to-openapi-yaml> <path-to-sdk-templates> <path-to-docs-templates>",
	Short: "Generate SDKs for the Cloud Firewall Manager API",
	Long:  `Generates SDKs for the Cloud Firewall Manager API using the OpenAPI specification.

The sdk-templates path should point to the directory containing the mustache templates for the SDK.
The docs-templates path should point to the directory where generated API docs will be copied (for Read the Docs).`,
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		openAPISpecPath := args[0]
		sdkTemplatesPath := args[1]
		docsTemplatesPath := args[2]
		if _, err := os.Stat(openAPISpecPath); os.IsNotExist(err) {
			pterm.Error.Printf("The specified OpenAPI spec file does not exist: %s\n", openAPISpecPath)
			os.Exit(1)
		}
		if _, err := os.Stat(sdkTemplatesPath); os.IsNotExist(err) {
			pterm.Error.Printf("The specified SDK templates directory does not exist: %s\n", sdkTemplatesPath)
			os.Exit(1)
		}
		if _, err := os.Stat(docsTemplatesPath); os.IsNotExist(err) {
			pterm.Error.Printf("The specified docs templates directory does not exist: %s\n", docsTemplatesPath)
			os.Exit(1)
		}
		data, err := os.ReadFile(openAPISpecPath)
		if err != nil {
			pterm.Error.Println("Failed to read OpenAPI spec file", err)
			os.Exit(1)
		}
		openApiSpec := models.OpenAPI{}
		if err = yaml.Unmarshal(data, &openApiSpec); err != nil {
			pterm.Error.Println("Failed to unmarshal OpenAPI spec", err)
			os.Exit(1)
		}
		useLocalInstallation, err := cmd.Flags().GetBool("use-local-installation")
		if err != nil {
			pterm.Error.Println("Failed to get publish flag", err)
			os.Exit(1)
		}

		version, err := generatePythonSdk(openAPISpecPath, sdkTemplatesPath, docsTemplatesPath, useLocalInstallation, &openApiSpec)
		if err != nil {
			pterm.Error.Println("Failed to generate Python SDK", err)
			os.Exit(1)
		}

		shouldPublish, err := cmd.Flags().GetBool("publish")
		if err != nil {
			pterm.Error.Println("Failed to get publish flag", err)
			os.Exit(1)
		}
		if shouldPublish {
			pypiToken, err := cmd.Flags().GetString("pypi-token")
			if err != nil {
				pterm.Error.Println("Failed to get PyPI token", err)
				os.Exit(1)
			}
			// golang cobra does not support nil flag values
			publishPythonSdk(pypiToken, *version)
		}
	},
}

func generatePythonSdk(openApiFile string, sdkTemplatesPath string, docsTemplatesPath string, useLocalInstallation bool, openApiSpec *models.OpenAPI) (*string, error) {
	// Generate Python SDK
	version, err := services.GetCurrentVersion(openApiSpec.Info.Version)
	if err != nil {
		pterm.Error.Println("Failed to get current version", err)
		return nil, err
	}
	pterm.Info.Printf("Generating Python SDK with version %s\n", *version)
	spinner, _ := pterm.DefaultSpinner.Start("Generating Python SDK...")
	err = services.GeneratePythonSdk(openApiFile, sdkTemplatesPath, *version, useLocalInstallation)
	if err != nil {
		spinner.Fail("Error generating Python SDK", err)
		os.Exit(1)
	}
	spinner.Success("Python SDK generated successfully")

	// Copy generated API docs to docs templates directory
	spinner, _ = pterm.DefaultSpinner.Start("Copying generated API docs...")
	err = services.BuildPythonSdkDocs(docsTemplatesPath)
	if err != nil {
		spinner.Fail("Error copying API docs", err)
		os.Exit(1)
	}
	spinner.Success("API docs copied successfully")

	return version, nil
}

func publishPythonSdk(pypiToken string, version string) {
	// Publish Python SDK
	spinner, _ := pterm.DefaultSpinner.Start("Publishing Python SDK...")
	var tokenPtr *string
	if pypiToken != "" {
		tokenPtr = &pypiToken
	}
	err := services.PublishPythonSdk(tokenPtr, version)
	if err != nil {
		spinner.Fail("Error publishing Python SDK", err)
		os.Exit(1)
	}
	spinner.Success("Python SDK published successfully")
}

func init() {
	rootCmd.AddCommand(generateSdksCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateSdksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	generateSdksCmd.Flags().BoolP("publish", "p", false, "Publish package to repositories after generation")
	generateSdksCmd.Flags().BoolP("use-local-installation", "l", false, "Install @openapitools/openapi-generator-cli locally and use local installation instead of global installation. This is required for Jenins agents, where the global installation fails.")
	generateSdksCmd.Flags().StringP("pypi-token", "t", "", "PyPI token to use for publishing the package. If not specified, it will pull the token from AWS secrets manager.")
}
