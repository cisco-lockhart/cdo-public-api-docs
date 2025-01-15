package services_test

import (
	"errors"
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
				} else {
					// validate args
					Expect(name).To(Equal("npx"))
					Expect(args).To(HaveLen(9))
					Expect(args).To(Equal([]string{"./node_modules/@openapitools/openapi-generator-cli", "generate", "-i", openapiFile, "-g", "python", "-o", "sdks/python", "--additional-properties=packageName=cdo-sdk-python,packageVersion=" + version}))

					output := "test output"
					return &MockCommandExecutor{output: &output}
				}
			}
			services.ExecCommand = mockShellCommandFunc
			err := services.GeneratePythonSdk(openapiFile, version, true)
			Expect(err).To(BeNil())
			Expect(execCount).To(Equal(2))
		})

		It("should generate a python SDK using a global installation", func() {
			openapiFile := "test-file"
			version := "1.0.0"
			mockShellCommandFunc := func(name string, args ...string) services.CommandExecutor {
				// validate args
				Expect(name).To(Equal("npx"))
				Expect(args).To(HaveLen(9))
				Expect(args).To(Equal([]string{"@openapitools/openapi-generator-cli", "generate", "-i", openapiFile, "-g", "python", "-o", "sdks/python", "--additional-properties=packageName=cdo-sdk-python,packageVersion=" + version}))

				output := "test output"
				return &MockCommandExecutor{output: &output}
			}
			services.ExecCommand = mockShellCommandFunc
			err := services.GeneratePythonSdk(openapiFile, version, false)
			Expect(err).To(BeNil())
		})

		It("should fail if python SDK generation fails", func() {
			openapiFile := "test-file"
			version := "1.0.0"
			expectedErrMsg := "test-error"
			mockShellCommandFunc := func(name string, args ...string) services.CommandExecutor {
				// validate args
				Expect(name).To(Equal("npx"))
				Expect(args).To(HaveLen(9))
				Expect(args).To(Equal([]string{"@openapitools/openapi-generator-cli", "generate", "-i", openapiFile, "-g", "python", "-o", "sdks/python", "--additional-properties=packageName=cdo-sdk-python,packageVersion=" + version}))

				return &MockCommandExecutor{err: errors.New(expectedErrMsg)}
			}
			services.ExecCommand = mockShellCommandFunc
			err := services.GeneratePythonSdk(openapiFile, version, false)
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(Equal(expectedErrMsg))
		})
	})
})
