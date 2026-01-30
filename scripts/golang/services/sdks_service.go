package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/pterm/pterm"
)

const region = "us-west-2"
const awsAccountId = "107042026245"
const secretName = "jenkins-pypi-credentials-mD4NdK"

func GeneratePythonSdk(openapiFile string, sdkTemplatesPath string, version string, useLocalInstallation bool) error {
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
		"-o", "sdks/python",
		"-t", sdkTemplatesPath,
		"--additional-properties=packageName=scc_firewall_manager_sdk,packageVersion="+version).CombinedOutput()
	if err != nil {
		return errors.New(fmt.Sprintf("Error: %s.\nSDK Output: %s", err.Error(), out))
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

func BuildPythonSdkDocs(docsTemplatePath string) error {
	// Copy generated API docs from sdks/python/docs/ to the docs template directory
	cmd := ExecCommand("bash", "-c", fmt.Sprintf("cp sdks/python/docs/*.md %s/docs/", docsTemplatePath))
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(fmt.Sprintf("Error copying generated docs: %s.\nOutput: %s", err.Error(), out))
	}

	// Generate mkdocs.yml with proper navigation
	err = generateMkdocsNav(docsTemplatePath)
	if err != nil {
		return err
	}

	pterm.Success.Printf("Generated API docs copied to %s/docs/\n", docsTemplatePath)
	return nil
}

func generateMkdocsNav(docsTemplatePath string) error {
	docsDir := fmt.Sprintf("%s/docs", docsTemplatePath)
	files, err := os.ReadDir(docsDir)
	if err != nil {
		return errors.New(fmt.Sprintf("Error reading docs directory: %s", err.Error()))
	}

	var apis []string
	var models []string

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".md") {
			continue
		}
		name := file.Name()
		if name == "index.md" || name == "getting-started.md" {
			continue
		}
		if strings.HasSuffix(name, "Api.md") {
			apis = append(apis, name)
		} else {
			models = append(models, name)
		}
	}

	sort.Strings(apis)
	sort.Strings(models)

	// Build mkdocs.yml content
	var sb strings.Builder
	sb.WriteString(`site_name: SCC Firewall Manager SDK
site_description: Python SDK for Cisco Security Cloud Control Firewall Manager API

theme:
  name: readthedocs

docs_dir: docs

plugins:
  - search

markdown_extensions:
  - toc:
      permalink: true

nav:
  - Home: index.md
  - Getting Started: getting-started.md
  - API Reference:
`)
	for _, api := range apis {
		displayName := strings.TrimSuffix(api, ".md")
		sb.WriteString(fmt.Sprintf("    - %s: %s\n", displayName, api))
	}
	sb.WriteString("  - Models:\n")
	for _, model := range models {
		displayName := strings.TrimSuffix(model, ".md")
		sb.WriteString(fmt.Sprintf("    - %s: %s\n", displayName, model))
	}

	mkdocsPath := fmt.Sprintf("%s/mkdocs.yml", docsTemplatePath)
	err = os.WriteFile(mkdocsPath, []byte(sb.String()), 0644)
	if err != nil {
		return errors.New(fmt.Sprintf("Error writing mkdocs.yml: %s", err.Error()))
	}

	return nil
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
