package services

func GeneratePostmanCollection(openapiFile string, postmanCollectionFile string) error {
	_, err := ExecCommand("npx", "openapi2postmanv2", "-s", openapiFile, "-o", postmanCollectionFile, "-O", "folderStrategy=Tags").Output()
	if err != nil {
		return err
	}
	return nil
}
