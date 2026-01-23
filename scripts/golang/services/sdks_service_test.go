package services_test

import (
	"errors"
	"fmt"

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
					Expect(args).To(Equal([]string{"./node_modules/@openapitools/openapi-generator-cli", "generate", "-i", openapiFile, "-g", "python", "-t", "openapi-templates", "-o", "sdks/python", "--additional-properties=packageName=scc_firewall_manager_sdk,packageVersion=" + version}))

					output := "test output"
					return &MockCommandExecutor{output: &output}
				} else {
					// validate args
					Expect(name).To(Equal("cp"))
					Expect(args).To(HaveLen(2))
					Expect(args).To(Equal([]string{openapiFile, "sdks/python/scc_firewall_manager_sdk/"}))

					output := "test output"
					return &MockCommandExecutor{output: &output}
				}
			}
			services.ExecCommand = mockShellCommandFunc
			err := services.GeneratePythonSdk(openapiFile, version, true)
			Expect(err).To(BeNil())
			Expect(execCount).To(Equal(3))
		})

		It("should generate a python SDK using a global installation", func() {
			openapiFile := "test-file"
			version := "1.0.0"
			execCount := 0
			mockShellCommandFunc := func(name string, args ...string) services.CommandExecutor {
				execCount++
				if execCount == 1 {
					// validate args
					Expect(name).To(Equal("npx"))
					Expect(args).To(HaveLen(11))
					Expect(args).To(Equal([]string{"@openapitools/openapi-generator-cli", "generate", "-i", openapiFile, "-g", "python", "-t", "openapi-templates", "-o", "sdks/python", "--additional-properties=packageName=scc_firewall_manager_sdk,packageVersion=" + version}))

					output := "test output"
					return &MockCommandExecutor{output: &output}
				} else {
					// validate args
					Expect(name).To(Equal("cp"))
					Expect(args).To(HaveLen(2))
					Expect(args).To(Equal([]string{openapiFile, "sdks/python/scc_firewall_manager_sdk/"}))
				}

				output := "test output"
				return &MockCommandExecutor{output: &output}
			}
			services.ExecCommand = mockShellCommandFunc
			err := services.GeneratePythonSdk(openapiFile, version, false)
			Expect(err).To(BeNil())
			Expect(execCount).To(Equal(2))
		})

		It("should fail if python SDK generation fails", func() {
			openapiFile := "test-file"
			version := "1.0.0"
			expectedSdkOutput := "test-output"
			errMsg := "test-error"
			expectedErr := errors.New(fmt.Sprintf("Error: %s.\nSDK Output: %s", errMsg, expectedSdkOutput))
			mockShellCommandFunc := func(name string, args ...string) services.CommandExecutor {
				// validate args
				Expect(name).To(Equal("npx"))
				Expect(args).To(HaveLen(11))
				Expect(args).To(Equal([]string{"@openapitools/openapi-generator-cli", "generate", "-i", openapiFile, "-g", "python", "-t", "openapi-templates", "-o", "sdks/python", "--additional-properties=packageName=scc_firewall_manager_sdk,packageVersion=" + version}))

				return &MockCommandExecutor{err: errors.New(errMsg), output: &expectedSdkOutput}
			}
			services.ExecCommand = mockShellCommandFunc
			err := services.GeneratePythonSdk(openapiFile, version, false)
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
})
