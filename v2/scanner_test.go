package v2_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	goparsec "github.com/bricef/goparsec"
)

var _ = Describe("Scanner", func() {
	It("Should end immediately when created with empty byte stream", func() {
		s := goparsec.NewScanner([]byte(``))
		Expect(s.Endof()).Should(BeTrue())
	})
})
