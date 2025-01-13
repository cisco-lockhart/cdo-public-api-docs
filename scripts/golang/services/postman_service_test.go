package services_test

import (
	"errors"
	"github.com/cisco-lockhart/cloud-fw-mgr-api-docs/services"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GeneratePostmanCollection", func() {
	var realExecCommand func(name string, args ...string) services.CommandExecutor
	BeforeEach(func() {
		realExecCommand = services.ExecCommand
	})

	AfterEach(func() {
		services.ExecCommand = realExecCommand
	})

	It("should call a command to generate a Postman collection", func() {
		openApiFile := "openapi.yaml"
		postmanCollection := "postman.json"
		mockShellCommandFunc := func(name string, args ...string) services.CommandExecutor {
			// validate args
			Expect(name).To(Equal("npx"))
			Expect(args).To(HaveLen(7))
			Expect(args).To(Equal([]string{"openapi2postmanv2", "-s", openApiFile, "-o", postmanCollection, "-O", "folderStrategy=Tags"}))

			output := "test output"
			return &MockCommandExecutor{output: &output}
		}
		services.ExecCommand = mockShellCommandFunc
		err := services.GeneratePostmanCollection(openApiFile, postmanCollection)
		Expect(err).To(BeNil())
	})

	It("should error out if the command errored out", func() {
		openApiFile := "openapi.yaml"
		postmanCollection := "postman.json"
		expectedErrMsg := "test-error"
		mockShellCommandFunc := func(name string, args ...string) services.CommandExecutor {
			// validate args
			return &MockCommandExecutor{output: nil, err: errors.New(expectedErrMsg)}
		}

		services.ExecCommand = mockShellCommandFunc
		err := services.GeneratePostmanCollection(openApiFile, postmanCollection)
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal(expectedErrMsg))
	})
})
