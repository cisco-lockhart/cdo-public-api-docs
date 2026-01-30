package services_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cisco-lockhart/cloud-fw-mgr-api-docs/services"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type MockSecretsManager struct {
	GetSecretFunc func(accountNumber string, region string, secretName string) (*string, error)
}

func (m *MockSecretsManager) GetSecret(accountNumber string, region string, secretName string) (*string, error) {
	return m.GetSecretFunc(accountNumber, region, secretName)
}

var _ = Describe("SdksService", func() {
	var realExecCommand func(name string, args ...string) services.CommandExecutor
	BeforeEach(func() {
		realExecCommand = services.ExecCommand
	})

	AfterEach(func() {
		services.ExecCommand = realExecCommand
	})

	Context("GeneratePythonSdk", func() {
		It("should generate a python SDK using a local installation", func() {
			openapiFile := "test-file"
			sdkTemplatesPath := "../../sdk-templates/python"
			openApiPackage := "@openapitools/openapi-generator-cli"
			version := "1.0.0"
			execCount := 0
			mockShellCommandFunc := func(name string, args ...string) services.CommandExecutor {
				execCount++
				if execCount == 1 {
					// validate args
					Expect(name).To(Equal("npm"))
					Expect(args).To(HaveLen(2))
					Expect(args).To(Equal([]string{"install", openApiPackage}))

					output := "test output"
					return &MockCommandExecutor{output: &output}
				} else if execCount == 2 {
					// validate args
					Expect(name).To(Equal("npx"))
					Expect(args).To(HaveLen(11))
					Expect(args).To(Equal([]string{"./node_modules/@openapitools/openapi-generator-cli", "generate", "-i", openapiFile, "-g", "python", "-o", "sdks/python", "-t", sdkTemplatesPath, "--additional-properties=packageName=scc_firewall_manager_sdk,packageVersion=" + version}))

					output := "test output"
					return &MockCommandExecutor{output: &output}
				}
			}
			services.ExecCommand = mockShellCommandFunc
			err := services.GeneratePythonSdk(openapiFile, sdkTemplatesPath, version, true)
			Expect(err).To(BeNil())
			Expect(execCount).To(Equal(3))
		})

		It("should generate a python SDK using a global installation", func() {
			openapiFile := "test-file"
			sdkTemplatesPath := "../../sdk-templates/python"
			version := "1.0.0"
			execCount := 0
			mockShellCommandFunc := func(name string, args ...string) services.CommandExecutor {
				// validate args
				Expect(name).To(Equal("npx"))
				Expect(args).To(HaveLen(11))
				Expect(args).To(Equal([]string{"@openapitools/openapi-generator-cli", "generate", "-i", openapiFile, "-g", "python", "-o", "sdks/python", "-t", sdkTemplatesPath, "--additional-properties=packageName=scc_firewall_manager_sdk,packageVersion=" + version}))

				output := "test output"
				return &MockCommandExecutor{output: &output}
			}
			services.ExecCommand = mockShellCommandFunc
			err := services.GeneratePythonSdk(openapiFile, sdkTemplatesPath, version, false)
			Expect(err).To(BeNil())
			Expect(execCount).To(Equal(2))
		})

		It("should fail if python SDK generation fails", func() {
			openapiFile := "test-file"
			sdkTemplatesPath := "../../sdk-templates/python"
			version := "1.0.0"
			expectedSdkOutput := "test-output"
			errMsg := "test-error"
			expectedErr := errors.New(fmt.Sprintf("Error: %s.\nSDK Output: %s", errMsg, expectedSdkOutput))
			mockShellCommandFunc := func(name string, args ...string) services.CommandExecutor {
				// validate args
				Expect(name).To(Equal("npx"))
				Expect(args).To(HaveLen(11))
				Expect(args).To(Equal([]string{"@openapitools/openapi-generator-cli", "generate", "-i", openapiFile, "-g", "python", "-o", "sdks/python", "-t", sdkTemplatesPath, "--additional-properties=packageName=scc_firewall_manager_sdk,packageVersion=" + version}))

				return &MockCommandExecutor{err: errors.New(errMsg), output: &expectedSdkOutput}
			}
			services.ExecCommand = mockShellCommandFunc
			err := services.GeneratePythonSdk(openapiFile, sdkTemplatesPath, version, false)
			Expect(err).ToNot(BeNil())
			Expect(err).To(Equal(expectedErr))
		})

		It("should fail if copying the OpenAPI spec fails", func() {
			openapiFile := "test-file"
			version := "1.0.0"
			generationOutput := "generation-output"
			copyOutput := "copy-output"
			copyErrMsg := "copy-error"
			expectedErr := errors.New(fmt.Sprintf("Error: %s.\nCopy OpenAPI Spec Output: %s", copyErrMsg, copyOutput))

			execCount := 0
			mockShellCommandFunc := func(name string, args ...string) services.CommandExecutor {
				execCount++
				if execCount == 1 {
					// successful SDK generation
					Expect(name).To(Equal("npx"))
					Expect(args).To(HaveLen(11))
					Expect(args).To(Equal([]string{"@openapitools/openapi-generator-cli", "generate", "-i", openapiFile, "-g", "python", "-t", "openapi-templates", "-o", "sdks/python", "--additional-properties=packageName=scc_firewall_manager_sdk,packageVersion=" + version}))

					return &MockCommandExecutor{output: &generationOutput}
				} else {
					// failing copy command
					Expect(name).To(Equal("cp"))
					Expect(args).To(HaveLen(2))
					Expect(args).To(Equal([]string{openapiFile, "sdks/python/scc_firewall_manager_sdk/"}))

					return &MockCommandExecutor{err: errors.New(copyErrMsg), output: &copyOutput}
				}
			}
			services.ExecCommand = mockShellCommandFunc
			err := services.GeneratePythonSdk(openapiFile, version, false)
			Expect(err).ToNot(BeNil())
			Expect(err).To(Equal(expectedErr))
			Expect(execCount).To(Equal(2))
		})
	})

	Context("BuildPythonSdkDocs", func() {
		var tempDir string

		BeforeEach(func() {
			var err error
			tempDir, err = os.MkdirTemp("", "sdk-docs-test")
			Expect(err).To(BeNil())

			// Create docs subdirectory
			err = os.MkdirAll(filepath.Join(tempDir, "docs"), 0755)
			Expect(err).To(BeNil())

			// Create test markdown files
			err = os.WriteFile(filepath.Join(tempDir, "docs", "index.md"), []byte("# Index"), 0644)
			Expect(err).To(BeNil())
			err = os.WriteFile(filepath.Join(tempDir, "docs", "getting-started.md"), []byte("# Getting Started"), 0644)
			Expect(err).To(BeNil())
			err = os.WriteFile(filepath.Join(tempDir, "docs", "TestApi.md"), []byte("# TestApi"), 0644)
			Expect(err).To(BeNil())
			err = os.WriteFile(filepath.Join(tempDir, "docs", "AnotherApi.md"), []byte("# AnotherApi"), 0644)
			Expect(err).To(BeNil())
			err = os.WriteFile(filepath.Join(tempDir, "docs", "TestModel.md"), []byte("# TestModel"), 0644)
			Expect(err).To(BeNil())
			err = os.WriteFile(filepath.Join(tempDir, "docs", "AnotherModel.md"), []byte("# AnotherModel"), 0644)
			Expect(err).To(BeNil())
		})

		AfterEach(func() {
			os.RemoveAll(tempDir)
		})

		It("should copy docs and generate mkdocs.yml with proper navigation", func() {
			mockShellCommandFunc := func(name string, args ...string) services.CommandExecutor {
				output := "test output"
				return &MockCommandExecutor{output: &output}
			}
			services.ExecCommand = mockShellCommandFunc

			err := services.BuildPythonSdkDocs(tempDir)
			Expect(err).To(BeNil())

			// Verify mkdocs.yml was created
			mkdocsPath := filepath.Join(tempDir, "mkdocs.yml")
			_, err = os.Stat(mkdocsPath)
			Expect(err).To(BeNil())

			// Read and verify mkdocs.yml content
			content, err := os.ReadFile(mkdocsPath)
			Expect(err).To(BeNil())
			mkdocsContent := string(content)

			// Verify structure
			Expect(mkdocsContent).To(ContainSubstring("site_name: SCC Firewall Manager SDK"))
			Expect(mkdocsContent).To(ContainSubstring("- Home: index.md"))
			Expect(mkdocsContent).To(ContainSubstring("- Getting Started: getting-started.md"))
			Expect(mkdocsContent).To(ContainSubstring("- API Reference:"))
			Expect(mkdocsContent).To(ContainSubstring("- AnotherApi: AnotherApi.md"))
			Expect(mkdocsContent).To(ContainSubstring("- TestApi: TestApi.md"))
			Expect(mkdocsContent).To(ContainSubstring("- Models:"))
			Expect(mkdocsContent).To(ContainSubstring("- AnotherModel: AnotherModel.md"))
			Expect(mkdocsContent).To(ContainSubstring("- TestModel: TestModel.md"))
		})

		It("should fail if copy command fails", func() {
			errMsg := "copy failed"
			mockShellCommandFunc := func(name string, args ...string) services.CommandExecutor {
				return &MockCommandExecutor{err: errors.New(errMsg), output: nil}
			}
			services.ExecCommand = mockShellCommandFunc

			err := services.BuildPythonSdkDocs(tempDir)
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(ContainSubstring("Error copying generated docs"))
		})

		It("should sort APIs and Models alphabetically", func() {
			mockShellCommandFunc := func(name string, args ...string) services.CommandExecutor {
				output := "test output"
				return &MockCommandExecutor{output: &output}
			}
			services.ExecCommand = mockShellCommandFunc

			err := services.BuildPythonSdkDocs(tempDir)
			Expect(err).To(BeNil())

			content, err := os.ReadFile(filepath.Join(tempDir, "mkdocs.yml"))
			Expect(err).To(BeNil())
			mkdocsContent := string(content)

			// AnotherApi should come before TestApi (alphabetical)
			anotherApiIndex := len(mkdocsContent) - len(mkdocsContent[findIndex(mkdocsContent, "AnotherApi"):])
			testApiIndex := len(mkdocsContent) - len(mkdocsContent[findIndex(mkdocsContent, "TestApi"):])
			Expect(anotherApiIndex).To(BeNumerically("<", testApiIndex))

			// AnotherModel should come before TestModel (alphabetical)
			anotherModelIndex := len(mkdocsContent) - len(mkdocsContent[findIndex(mkdocsContent, "AnotherModel"):])
			testModelIndex := len(mkdocsContent) - len(mkdocsContent[findIndex(mkdocsContent, "TestModel"):])
			Expect(anotherModelIndex).To(BeNumerically("<", testModelIndex))
		})
	})
})

func findIndex(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
