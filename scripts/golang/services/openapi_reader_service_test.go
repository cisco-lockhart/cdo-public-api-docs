package services_test

import (
	"github.com/cisco-lockhart/cloud-fw-mgr-api-docs/services"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("LoadOpenApi", func() {
	It("Should successfully load an OpenAPI spec", func() {
		openApi, err := services.LoadOpenApi("https://edge.us.cdo.cisco.com/api/platform/public-api/v3/api-docs.yaml")
		Expect(openApi).NotTo(BeNil())
		Expect(err).To(BeNil())
	})

	It("Should fail to load an OpenAPI spec if the URL does not exist", func() {
		// this URL does not exist
		var url = "https://cisco.doesnotexist.com/v3/openapi.yaml"
		openApi, err := services.LoadOpenApi(url)
		Expect(openApi).To(BeNil())
		Expect(err).NotTo(BeNil())
	})

	It("Should fail to load an OpenAPI spec if the URL is not an OpenAPI YAML", func() {
		var url = "https://www.cisco.com"
		openApi, err := services.LoadOpenApi(url)
		Expect(openApi).To(BeNil())
		Expect(err).NotTo(BeNil())
	})
})
