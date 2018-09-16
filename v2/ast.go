package v2

type Source struct {
	File  string
	Line  int
	Index int
}

type Node interface {
	Name() Identifier
	Children() []Node
	Segment() string
	Source() Source
	IsTerminal() bool
}

func NewTerminal(name Identifier, token string, cursor Cursor) Node {
	return nil
}

func NewNonTerminal() Node {
	return nil
}

type NodeTransformer func(Node) Node

type NodePredicate func(Node) bool

func TransformAll(root Node, transform NodeTransformer) Node {
	return nil
}

func TransformMatch(root Node, predicate NodePredicate, transform NodeTransformer) Node {
	return nil
}

func TransformMatches(root Node, matchers []struct {
	NodePredicate
	NodeTransformer
}) Node {
	return nil
}
