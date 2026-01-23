package services

import (
	"context"
	"errors"
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
	out, err := ExecCommand("npx",
		commandName,
		"generate",
		"-i", openapiFile,
		"-g", "python",
		"-t", "openapi-templates/python",
		"-o", "sdks/python",
		"--additional-properties=packageName=scc_firewall_manager_sdk,packageVersion="+version).CombinedOutput()
	if err != nil {
		return errors.New(fmt.Sprintf("Error: %s.\nSDK Output: %s", err.Error(), out))
	}

	out, err = ExecCommand("cp", openapiFile, "sdks/python/scc_firewall_manager_sdk/").CombinedOutput()
	if err != nil {
		return errors.New(fmt.Sprintf("Error: %s.\nCopy OpenAPI Spec Output: %s", err.Error(), out))
	}
	
	return nil
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
	out, err := cmd.Output()
	if err != nil {
		pterm.Error.Println(out)
		pterm.Error.Println(err)
	}

	cmd = ExecCommand("twine", "upload", "-u", "__token__", "-p", *pypiToken, fmt.Sprintf("sdks/python/dist/scc_firewall_manager_sdk-%s-py3-none-any.whl", version))
	out, err = cmd.Output()
	if err != nil {
		pterm.Error.Println(string(out))
		pterm.Error.Println(err.Error())
	}

	return err
}

func getPyPiTokenFromSecretsManager() (*string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
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
