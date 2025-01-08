/*
Copyright © 2025 Cisco Systems
*/
package cmd

import (
	"github.com/cisco-lockhart/fcm-api-docs-generator/services"
	"github.com/pterm/pterm"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var nodePackagesToInstall = []string{"openapi-to-postmanv2", "@openapitools/openapi-generator-cli"}

// installPreReqsCmd represents the installPreReqs command
var installPreReqsCmd = &cobra.Command{
	Use:   "install-pre-reqs",
	Short: "Install pre-requisites for the CLI tool",
	Long:  `Installs all the pre-requisites (such as NPM etc) to allow the CLI tool to generate SDKs etc. Only supports MacOS at the moment. Requires npm to be installed and in your path.`,
	Run: func(cmd *cobra.Command, args []string) {
		if runtime.GOOS != "darwin" {
			pterm.Error.Println("This command is only supported on MacOS at the moment.")
			return
		}
		npmVersionCmd := exec.Command("npm", "-v")
		_, err := npmVersionCmd.Output()
		if err != nil {
			pterm.Error.Println("npm is not installed.")
			return
		}
		for _, nodePackage := range nodePackagesToInstall {
			spinner, _ := pterm.DefaultSpinner.Start("Installing node package " + nodePackage)
			err := services.InstallNodePackage(nodePackage)
			if err != nil {
				spinner.Fail("Error installing node package", nodePackage)
				return
			}
			spinner.Success("Node package " + nodePackage + " installed successfully")
		}

		pterm.Success.Println("All pre-requisites installed successfully.")
	},
}

func init() {
	rootCmd.AddCommand(installPreReqsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installPreReqsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installPreReqsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
