package services_test

import (
	"github.com/cisco-lockhart/cloud-fw-mgr-api-docs/services"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ConfigReaderService", func() {
	Describe("LoadConfig", func() {
		It("should return a valid configuration", func() {
			url := "https://raw.githubusercontent.com/cisco-lockhart/cdo-public-api-docs/refs/heads/LH-92055-add-cdfmc-docs-to-cdo-docs/cloud-fw-mgr-api-docs.config.yaml"
			config, err := services.LoadConfig(url)
			Expect(config).NotTo(BeNil())
			Expect(err).To(BeNil())
			Expect(config.Services).NotTo(BeEmpty())
		})

		It("should fail if the URL is not a valid YAML", func() {
			invalidUrl := "https://raw.githubusercontent.com/cisco-lockhart/cdo-public-api-docs/refs/heads/main/README.md"
			config, err := services.LoadConfig(invalidUrl)
			Expect(config).To(BeNil())
			Expect(err).NotTo(BeNil())
		})
	})
})
