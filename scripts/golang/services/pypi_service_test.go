package services_test

import (
	"github.com/cisco-lockhart/cloud-fw-mgr-api-docs/services"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetCurrentVersion", func() {
	It("should return the current version", func() {
		version, err := services.GetCurrentVersion("0.0.1")
		Expect(version).NotTo(BeNil())
		Expect(err).To(BeNil())
	})
})
