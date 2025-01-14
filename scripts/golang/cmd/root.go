/*
Copyright © 2025 Cisco Systems
*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
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
	rootCmd.PersistentFlags().StringP("config", "c", "https://raw.githubusercontent.com/cisco-lockhart/cdo-public-api-docs/refs/heads/main/cloud-fw-mgr-api-docs.config.yaml", "config URL")
}
