package v2

import (
	parsec "github.com/bricef/goparsec"
)

// Parser function parses input text encapsulated by Scanner, higher
// order parsers are constructed using combinators.
type Parser func(parsec.Scanner) (ASTNode, parsec.Scanner)

//
// Utility types
//

type Identifier string
type Regex string
type ShortCircuitCallback func(ASTNode, parsec.Scanner) ASTNode

//
// Parser Constructors
//

func And(name Identifier, parsers ...interface{}) Parser {}

func Or(name Identifier, parsers ...interface{}) Parser {}

func Kleen(name Identifier, parser interface{}) Parser {}

func Many(name Identifier, parser interface{}) Parser {}

func ManyUntil(name Identifier, itemParser interface{}, stopParser interface{}) Parser {}

func Maybe(name Identifier, parser interface{}) Parser {}

func ShortCircuit(parser Parser, callback ShortCircuitCallback) Parser {}

func Char(name Identifier) Parser {}

func float(name Identifier) Parser {}

func Hex(name Identifier) Parser {}

func Int(name Identifier) Parser {}

func Oct(name Identifier) Parser {}

func String(name Identifier) Parser {}

func Ident(name Identifier) Parser {}

func Atom(name Identifier, atom string) Parser {}

func AtomExact(name Identifier, atom string) Parser {}

func Token(name Identifier, r Regex) Parser {}

func TokenExact(name Identifier, r Regex) Parser {}

func OrToken(name Identifier, r Regex) Parser {}

func End() Parser {}

func NotEnd() Parser {}

//
// AST constructiona and manipulation
//

type Source struct {
	File  string
	Line  int
	Index int
}

type ASTNode interface {
	Name() Identifier
	Children() []ASTNode
	Segment() string
	Source() Source
	IsTerminal() bool
}

type NodeTransformer func(ASTNode) ASTNode

type NodePredicate func(ASTNode) bool

func TransformAll(root ASTNode, transform NodeTransformer) ASTNode {}

func TransformMatch(root ASTNode, predicate NodePredicate, transform NodeTransformer) ASTNode {}

func TransformMatches(root ASTNode, matchers []struct {
	NodePredicate
	NodeTransformer
}) ASTNode {
	return nil
}

//
// Generating output using ASTs
//

type Output interface{}

type NodeGenerator func(ASTNode, []Output) Output

func GenerateAll(root ASTNode, generator NodeGenerator) Output {}

func FoldUp(root ASTNode, generators []struct {
	NodePredicate
	NodeGenerator
}) Output {
}

func IsA(name Identifier) NodePredicate {
	// IsA("FORLOOP")(node) -> T/F
	return func(n ASTNode) {
		return n.Name() == name
	}
}

func Map(nodes []ASTNode, generator NodeGenerator) []Output {
	out := make([]Out, len(nodes))
	for i, node := range nodes {
		out[i] = generator(node, nil)
	}
	return out
}

//
// Querying the syntax tree dynamically
//
func QueryFilter(root ASTNode, predicate NodePredicate) []ASTNode {}

func QueryTerminals(root ASTNode) []ASTNode {
	return QueryFilter(root, func(n ASTNode) {
		return n.IsTerminal()
	})
}

type QueryString string

func QueryOne(root ASTNode, query QueryString) ASTNode {}

func QueryMany(root ASTNode, query QueryString) []ASTNode {}

func Walk(root ASTNode) chan ASTNode {}
