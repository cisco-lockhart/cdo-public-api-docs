/*
Copyright © 2025 Cisco Systems
*/
package cmd

import (
	"fmt"
	"github.com/cisco-lockhart/cloud-fw-mgr-api-docs/services"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var generatePostmanCollectionCmd = &cobra.Command{
	Use:   "generate-postman-collection",
	Short: "Generate a Postman collection from an OpenAPI spec",
	Run: func(cmd *cobra.Command, args []string) {
		spinner, _ := pterm.DefaultSpinner.Start("Generating Postman collection...")
		openapiFile, err := cmd.Flags().GetString("openapi-file")
		if err != nil {
			spinner.Fail("Failed to get OpenAPI file", err)
			os.Exit(1)
		}
		postmanCollectionFile, err := cmd.Flags().GetString("postman-collection-file")
		if err != nil {
			spinner.Fail("Failed to get Postman collection file", err)
			os.Exit(1)
		}
		err = services.GeneratePostmanCollection(openapiFile, postmanCollectionFile)
		if err != nil {
			spinner.Fail("Failed to generate Postman collection", err)
			os.Exit(1)
		}
		spinner.Success(fmt.Sprintf("Postman collection generated successfully at %s", postmanCollectionFile))
	},
}

func init() {
	rootCmd.AddCommand(generatePostmanCollectionCmd)
	currDir, err := os.Getwd()
	if err != nil {
		pterm.Error.Println("Failed to get current working directory", err)
	}

	defaultInputFile := filepath.Join(currDir, "openapi.yaml")
	defaultOutputFile := filepath.Join(currDir, "postman_collection.json")
	generatePostmanCollectionCmd.Flags().StringP("openapi-file", "i", defaultInputFile, fmt.Sprintf("Specify the OpenAPI spec file to convert to a Postman collection", defaultInputFile))
	generatePostmanCollectionCmd.Flags().StringP("postman-collection-file", "o", defaultOutputFile, fmt.Sprintf("Specify the path to write the Postman collection to", defaultOutputFile))
}
