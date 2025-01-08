package services

import "os/exec"

func GeneratePythonSdk(openapiFile string, version string) error {
	_, err := exec.Command("npx",
		"@openapitools/openapi-generator-cli",
		"generate",
		"-i", openapiFile,
		"-g", "python",
		"-o", "sdks/python",
		"--additional-properties=packageName=cdo-sdk-python,packageVersion="+version).Output()
	return err
}

func PublishPythonSdk(pypiToken string) error {
	if pypiToken == "" {
		return nil
	}
	return nil
}
