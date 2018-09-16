package v2_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	gp "./"
	goparsec "github.com/bricef/goparsec"
)

var _ = Describe("V2", func() {
	Describe("Parsers", func() {
		Describe("Token", func() {
			It("should ignore a scanner when initialised with the empty string.", func() {
				p := gp.Token("", "")
				s := goparsec.NewScanner([]byte(``))

				node, newScanner := p(s)

				Expect(node).ShouldNot(HaveOccurred())
				Expect(newScanner).Should(Equal(s))
			})

			It("should match a token when present in the scanner", func() {
				p := gp.Token("", "a")
				s := goparsec.NewScanner([]byte(`a`))

				node, newScanner := p(s)

				Expect(node).ShouldNot(BeNil())
				Expect(node).Should(Equal(gp.NewTerminal("", "a", 0)))
				Expect(newScanner).ShouldNot(Equal(s))
			})
		})
	})
})
