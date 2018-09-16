package v2_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	gp "./"
	goparsec "github.com/bricef/goparsec"
)

var _ = Describe("V2", func() {
	Describe("Scanner", func() {
		It("Should end immediately when created with empty byte stream", func() {
			s := goparsec.NewScanner([]byte(``))
			Expect(s.Endof()).Should(BeTrue())
		})
	})

	Describe("Parsers", func() {
		Describe("Token", func() {
			It("should ignore a scanner when initialised with the empty string.", func() {
				p := gp.Token("", "")
				s := goparsec.NewScanner([]byte(``))

				node, newScanner := p(s)

				Expect(node).ShouldNot(HaveOccurred())
				Expect(newScanner).Should(Equal(s))
			})
		})
	})
})
