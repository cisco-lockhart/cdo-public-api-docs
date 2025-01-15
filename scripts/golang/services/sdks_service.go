package services

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/pterm/pterm"
)

const region = "us-west-2"
const awsAccountId = "107042026245"
const secretName = "jenkins-pypi-credentials-mD4NdK"

func GeneratePythonSdk(openapiFile string, version string, useLocalInstallation bool) error {
	var commandName string
	if useLocalInstallation {
		err := InstallNodePackageGloballyOrLocally("@openapitools/openapi-generator-cli", false)
		if err != nil {
			return err
		}
		commandName = "./node_modules/@openapitools/openapi-generator-cli"
	} else {
		commandName = "@openapitools/openapi-generator-cli"
	}
	_, err := ExecCommand("npx",
		commandName,
		"generate",
		"-i", openapiFile,
		"-g", "python",
		"-o", "sdks/python",
		"--additional-properties=packageName=cdo-sdk-python,packageVersion="+version).Output()
	return err
}

func PublishPythonSdk(pypiToken *string, version string) error {
	if pypiToken == nil {
		var err error
		pypiToken, err = getPyPiTokenFromSecretsManager()
		if err != nil {
			return err
		}
	}
	cmd := ExecCommand("bash", "-c", "cd sdks/python && rm -rf dist build *.egg-info && python3 setup.py sdist bdist_wheel")
	_, err := cmd.Output()
	if err != nil {
		pterm.Error.Println(err)
	}

	cmd = ExecCommand("twine", "upload", "-u", "__token__", "-p", *pypiToken, fmt.Sprintf("sdks/python/dist/cdo_sdk_python-%s-py3-none-any.whl", version))
	_, err = cmd.Output()
	if err != nil {
		pterm.Error.Println(err)
	}

	return err
}

func getPyPiTokenFromSecretsManager() (*string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		return nil, err
	}
	secretsManagerClient := secretsmanager.NewFromConfig(cfg)
	secretId := fmt.Sprintf("arn:aws:secretsmanager:%s:%s:secret:%s", region, awsAccountId, secretName)
	secretValue, err := secretsManagerClient.GetSecretValue(context.TODO(), &secretsmanager.GetSecretValueInput{
		SecretId: &secretId,
	})
	if err != nil {
		return nil, err
	}

	return secretValue.SecretString, nil
}
