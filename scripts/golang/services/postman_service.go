package services

import "os/exec"

func GeneratePostmanCollection(openapiFile string, postmanCollectionFile string) error {
	_, err := exec.Command("npx", "openapi2postmanv2", "-s", openapiFile, "-o", postmanCollectionFile, "-O", "folderStrategy=Tags").Output()
	if err != nil {
		return err
	}
	return nil
}
