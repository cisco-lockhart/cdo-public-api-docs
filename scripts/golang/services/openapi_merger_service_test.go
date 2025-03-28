package services_test

import (
	"github.com/cisco-lockhart/cloud-fw-mgr-api-docs/models"
	"github.com/cisco-lockhart/cloud-fw-mgr-api-docs/services"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("MergeOpenApiSpecs", func() {
	It("should merge multiple OpenAPI specs", func() {
		prefixToAddForService1 := "v1/another-custom-prefix/"
		prefixToAddForService2 := "/custom-prefix"
		config := models.Config{
			Info: models.Info{
				Cdo: models.InfoData{
					Title:       "API Docs",
					Version:     "1.0.0",
					Description: "Description",
					Contact: models.Contact{
						Name:  "SCC TAC",
						Email: "cdo.tac@cisco.com",
					},
				},
			},
			Servers: []models.Server{models.Server{
				URL:         "https://edge.us.cdo.cisco.com",
				Description: "US",
			}},
			Services: []models.Service{
				{
					Name:        "service1",
					URL:         "https://service1.us.cdo.cisco.com",
					PrefixToAdd: &prefixToAddForService1,
				}, {
					Name:        "service2",
					URL:         "https://service2.us.cdo.cisco.com",
					PrefixToAdd: &prefixToAddForService2,
				}},
		}

		service1Paths := map[string]interface{}{
			"/path1": map[string]interface{}{},
			"/path2": map[string]interface{}{},
		}
		service1Schemas := map[string]interface{}{
			"schema1": map[string]interface{}{},
			"schema2": map[string]interface{}{},
		}
		service1Responses := map[string]interface{}{
			"response1": map[string]interface{}{},
			"response2": map[string]interface{}{},
		}
		service2Paths := map[string]interface{}{
			"/path3": map[string]interface{}{},
			"/path4": map[string]interface{}{},
		}
		service2Schemas := map[string]interface{}{
			"schema3": map[string]interface{}{},
			"schema2": map[string]interface{}{},
		}
		service2Responses := map[string]interface{}{
			"response3": map[string]interface{}{},
			"response2": map[string]interface{}{},
		}

		serviceSpecs := map[string]*models.OpenAPI{
			"service1": {
				OpenAPI: "",
				Paths:   service1Paths,
				Components: models.Components{
					Schemas:   service1Schemas,
					Responses: service1Responses,
				},
				Definitions: nil,
			},
			"service2": {
				OpenAPI: "",
				Paths:   service2Paths,
				Components: models.Components{
					Schemas:   service2Schemas,
					Responses: service2Responses,
				},
				Definitions: nil,
			},
		}

		mergedSpec, err := services.MergeOpenApiSpecs(serviceSpecs, &config)
		Expect(err).To(BeNil())
		Expect(mergedSpec).NotTo(BeNil())
		Expect(mergedSpec.Paths).To(HaveLen(4))
		Expect(mergedSpec.Components.Schemas).To(HaveLen(3))
		Expect(mergedSpec.Components.Responses).To(HaveLen(3))
		// Verify that /path3 and /path4 are labeled /custom-prefix/path3 and /custom-prefix/path4
		Expect(mergedSpec.Paths).To(HaveKey("/v1/another-custom-prefix/path1"))
		Expect(mergedSpec.Paths).To(HaveKey("/v1/another-custom-prefix/path2"))
		Expect(mergedSpec.Paths).To(HaveKey("/custom-prefix/path3"))
		Expect(mergedSpec.Paths).To(HaveKey("/custom-prefix/path4"))
	})

	It("Should fail on conflicting paths", func() {
		config := models.Config{
			Info: models.Info{
				Cdo: models.InfoData{
					Title:       "API Docs",
					Version:     "1.0.0",
					Description: "Description",
					Contact: models.Contact{
						Name:  "SCC TAC",
						Email: "cdo.tac@cisco.com",
					},
				},
			},
			Servers: []models.Server{models.Server{
				URL:         "https://edge.us.cdo.cisco.com",
				Description: "US",
			}},
		}

		service1Paths := map[string]interface{}{
			"/path1": map[string]interface{}{},
			"/path2": map[string]interface{}{},
		}
		service2Paths := map[string]interface{}{
			"/path1": map[string]interface{}{},
			"/path4": map[string]interface{}{},
		}

		serviceSpecs := map[string]*models.OpenAPI{
			"service1": {
				OpenAPI:     "",
				Paths:       service1Paths,
				Definitions: nil,
			},
			"service2": {
				OpenAPI:     "",
				Paths:       service2Paths,
				Definitions: nil,
			},
		}

		mergedSpec, err := services.MergeOpenApiSpecs(serviceSpecs, &config)
		Expect(err).NotTo(BeNil())
		Expect(mergedSpec).To(BeNil())
	})
})
