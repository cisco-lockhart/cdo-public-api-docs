package services_test

import (
	"errors"
	"github.com/cisco-lockhart/cloud-fw-mgr-api-docs/services"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type MockCommandExecutor struct {
	// Used to stub the return of the Output method
	// Could add other properties depending on testing needs
	output *string
	err    error
}

// Implements the commandExecutor interface
func (m *MockCommandExecutor) Output() ([]byte, error) {
	if m.output == nil {
		return nil, m.err
	} else {
		return []byte(*m.output), m.err
	}
}

var _ = Describe("InstallerService", func() {
	Context("InstallNodePackage", func() {
		var realExecCommand func(name string, args ...string) services.CommandExecutor
		BeforeEach(func() {
			realExecCommand = services.ExecCommand
		})

		AfterEach(func() {
			services.ExecCommand = realExecCommand
		})

		It("should install a node package globally", func() {
			command := "test-pkg"
			mockShellCommandFunc := func(name string, args ...string) services.CommandExecutor {
				// validate args
				Expect(name).To(Equal("npm"))
				Expect(args).To(HaveLen(3))
				Expect(args).To(Equal([]string{"install", "-g", command}))

				output := "test output"
				return &MockCommandExecutor{output: &output}
			}
			services.ExecCommand = mockShellCommandFunc
			err := services.InstallNodePackage(command)
			Expect(err).To(BeNil())
		})

		It("should fail if node package installation fails", func() {
			command := "test-pkg"
			expectedErrMsg := "test-error"
			mockShellCommandFunc := func(name string, args ...string) services.CommandExecutor {
				// validate args
				Expect(name).To(Equal("npm"))
				Expect(args).To(HaveLen(3))
				Expect(args).To(Equal([]string{"install", "-g", command}))

				return &MockCommandExecutor{output: nil, err: errors.New(expectedErrMsg)}
			}
			services.ExecCommand = mockShellCommandFunc
			err := services.InstallNodePackage(command)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(expectedErrMsg))
		})

		It("should install a node package locally", func() {
			command := "test-pkg"
			mockShellCommandFunc := func(name string, args ...string) services.CommandExecutor {
				// validate args
				Expect(name).To(Equal("npm"))
				Expect(args).To(HaveLen(2))
				Expect(args).To(Equal([]string{"install", command}))

				output := "test output"
				return &MockCommandExecutor{output: &output}
			}
			services.ExecCommand = mockShellCommandFunc
			err := services.InstallNodePackageGloballyOrLocally(command, false)
			Expect(err).To(BeNil())
		})
	})

	Context("InstallHomebrewPackage", func() {
		It("should install a homebrew package", func() {
			command := "test-pkg"
			mockShellCommandFunc := func(name string, args ...string) services.CommandExecutor {
				// validate args
				Expect(name).To(Equal("brew"))
				Expect(args).To(HaveLen(2))
				Expect(args).To(Equal([]string{"install", command}))

				output := "test output"
				return &MockCommandExecutor{output: &output}
			}
			services.ExecCommand = mockShellCommandFunc
			err := services.InstallHomebrewPackage(command)
			Expect(err).To(BeNil())
		})

		It("should fail to install a Homebrew package", func() {
			command := "test-pkg"
			expectedErrMsg := "test-error"
			mockShellCommandFunc := func(name string, args ...string) services.CommandExecutor {
				// validate args
				Expect(name).To(Equal("brew"))
				Expect(args).To(HaveLen(2))
				Expect(args).To(Equal([]string{"install", command}))

				return &MockCommandExecutor{output: nil, err: errors.New(expectedErrMsg)}
			}
			services.ExecCommand = mockShellCommandFunc
			err := services.InstallHomebrewPackage(command)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(expectedErrMsg))
		})
	})

	Context("InstallPipxPackage", func() {
		It("should install Pipx package", func() {
			command := "test-pkg"
			mockShellCommandFunc := func(name string, args ...string) services.CommandExecutor {
				// validate args
				Expect(name).To(Equal("pipx"))
				Expect(args).To(HaveLen(3))
				Expect(args).To(Equal([]string{"install", command, "--include-deps"}))

				output := "test output"
				return &MockCommandExecutor{output: &output}
			}
			services.ExecCommand = mockShellCommandFunc
			err := services.InstallPipxPackage(command)
			Expect(err).To(BeNil())
		})

		It("should fail to install a Pipx package", func() {
			command := "test-pkg"
			expectedErrMsg := "test-error"
			mockShellCommandFunc := func(name string, args ...string) services.CommandExecutor {
				// validate args
				Expect(name).To(Equal("pipx"))
				Expect(args).To(HaveLen(3))
				Expect(args).To(Equal([]string{"install", command, "--include-deps"}))

				return &MockCommandExecutor{output: nil, err: errors.New(expectedErrMsg)}
			}
			services.ExecCommand = mockShellCommandFunc
			err := services.InstallPipxPackage(command)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(expectedErrMsg))
		})
	})
})
