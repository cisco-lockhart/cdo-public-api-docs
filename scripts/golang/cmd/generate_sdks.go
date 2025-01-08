/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/cisco-lockhart/fcm-api-docs-generator/models"
	"github.com/cisco-lockhart/fcm-api-docs-generator/services"
	"github.com/pterm/pterm"
	"gopkg.in/yaml.v3"
	"os"

	"github.com/spf13/cobra"
)

// generateSdksCmd represents the generateSdks command
var generateSdksCmd = &cobra.Command{
	Use:   "generate-sdks <path-to-openapi-yaml>",
	Short: "Generate SDKs for the Cloud Firewall Manager API",
	Long:  `Generates SDKs for the Cloud Firewall Manager API using the OpenAPI specification.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		openAPISpecPath := args[0]
		if _, err := os.Stat(openAPISpecPath); os.IsNotExist(err) {
			pterm.Error.Printf("The specified OpenAPI spec file does not exist: %s\n", openAPISpecPath)
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

		generatePythonSdk(openAPISpecPath, &openApiSpec)

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
			publishPythonSdk(pypiToken)
		}
	},
}

func generatePythonSdk(openApiFile string, openApiSpec *models.OpenAPI) {
	// Generate Python SDK
	version, err := services.GetCurrentVersion(openApiSpec.Info.Version)
	if err != nil {
		pterm.Error.Println("Failed to get current version", err)
		return
	}
	pterm.Info.Printf("Generating Python SDK with version %s\n", *version)
	spinner, _ := pterm.DefaultSpinner.Start("Generating Python SDK...")
	err = services.GeneratePythonSdk(openApiFile, *version)
	if err != nil {
		spinner.Fail("Error generating Python SDK")
		os.Exit(1)
	}
	spinner.Success("Python SDK generated successfully")
}

func publishPythonSdk(pypiToken string) {
	// Publish Python SDK
	spinner, _ := pterm.DefaultSpinner.Start("Publishing Python SDK...")
	err := services.PublishPythonSdk(pypiToken)
	if err != nil {
		spinner.Fail("Error publishing Python SDK")
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
	generateSdksCmd.Flags().StringP("pypi-token", "t", "", "PyPI token to use for publishing the package. If not specified, it will pull the token from AWS secrets manager.")
}
