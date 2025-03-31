package test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"latentlab.cc/easyhousing/api"
)

var _ = Describe("Sample Integration Test", Label(IntegrationTestLabel), func() {
	It("Doesn't throw an error", func() {
		message := "Hello World"
		res, err := Client.Echo(ctx, &api.EchoRequest{
			Message: &message,
		})

		Expect(err).ToNot(HaveOccurred())
		Expect(*res.Message).To(Equal(message))
	})
})
