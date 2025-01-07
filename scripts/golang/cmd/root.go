/*
Copyright © 2025 Cisco Systems
*/
package cmd

import (
	"fmt"
	"github.com/pterm/pterm"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cloud-fw-mgr-api-docs",
	Short: "CLI tool to generate API documentation and SDKs for Cloud Firewall Manager",
	Long:  `This tool provides commands to generate API documentation and SDKs for Cloud Firewall Manager.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		pterm.Error.Println("Failed to get home directory", err)
	}
	defaultConfigFilePath := filepath.Join(homeDir, ".cloud-fw-mgr-api-docs.config.yaml")
	rootCmd.PersistentFlags().StringP("config", "c", defaultConfigFilePath, fmt.Sprintf("config file (default is %s)", defaultConfigFilePath))
}
