package v2_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	gp "./"
	goparsec "github.com/bricef/goparsec"
)

var _ = Describe("Parsers", func() {
	Describe("Token", func() {
		It("should ignore a scanner when initialised with an empty pattern.", func() {
			p := gp.Token("", "")
			s := goparsec.NewScanner([]byte(``))

			node, newScanner := p(s)

			Expect(node).ShouldNot(HaveOccurred())
			Expect(newScanner).Should(Equal(s))
		})

		It("should match a token when present in the scanner and create a terminal.", func() {
			p := gp.Token("", "a")
			s := goparsec.NewScanner([]byte(`a`))

			node, newScanner := p(s)

			Expect(node).ShouldNot(BeNil())
			Expect(node).Should(Equal(gp.NewTerminal("", "a", gp.Source{"", 0, 0, 0})))
			Expect(newScanner).ShouldNot(Equal(s))
		})

		It("Should accept a name and create nodes with that name.", func() {
			name := "SomeRandomName"
			p := gp.Token(name, "a")
			s := goparsec.NewScanner([]byte(`a`))

			node, _ := p(s)

			Expect(node.Name()).Should(Equal(name))
		})

		It("Should fail to parse a stream that doesn't match the pattern.", func() {
			p := gp.Token("", "a")
			s := goparsec.NewScanner([]byte(`b`))

			node, newScanner := p(s)

			Expect(node).Should(BeNil())
			Expect(newScanner).Should(Equal(s))
		})

		It("Should ignore leading spaces in the stream", func() {
			p := gp.Token("", "a")
			s := goparsec.NewScanner([]byte(`     a`))

			node, _ := p(s)

			Expect(node).ShouldNot(BeNil())
			Expect(node.Value()).Should(Equal("a"))
		})

		It("Should correctly set the position of the token ignoring leading spaces in the stream", func() {
			p := gp.Token("", "a")
			s := goparsec.NewScanner([]byte(`     a`))

			node, _ := p(s)

			Expect(node).ShouldNot(BeNil())
			Expect(node.Source().Cursor).Should(Equal(5))
		})
	})
})
