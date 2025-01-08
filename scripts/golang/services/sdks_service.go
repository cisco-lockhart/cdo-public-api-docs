package services

import "os/exec"

func GeneratePythonSdk(openapiFile string, version string) error {
	_, err := exec.Command("npx",
		"@openapitools/openapi-generator-cli",
		"generate",
		"-i", openapiFile,
		"-g", "python",
		"-o", "sdks/python",
		"--additional-properties=packageName=cloud_firewall_manager,packageVersion="+version).Output()
	return err
}
