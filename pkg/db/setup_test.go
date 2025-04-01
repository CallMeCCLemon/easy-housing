package db

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Setup", func() {
	It("Can setup the database", func() {
		Expect(Setup()).ToNot(HaveOccurred())
	})
})
