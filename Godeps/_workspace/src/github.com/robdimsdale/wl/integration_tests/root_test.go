package wl_integration_test

import (
	"github.com/justincampbell/mylist/Godeps/_workspace/src/github.com/robdimsdale/wl"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("basic root functionality", func() {
	It("gets root correctly", func() {
		var err error
		var root wl.Root
		Eventually(func() error {
			root, err = client.Root()
			return err
		}).Should(Succeed())

		Expect(root.ID).To(BeNumerically(">", 0))
	})
})
