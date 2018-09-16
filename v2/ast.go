package v2

type Source struct {
	File  string
	Line  int
	Index Cursor
}

type Node interface {
	Name() Identifier
	Children() []Node
	Segment() string
	Source() Source
	IsTerminal() bool
}

//
// Terminal
//
type terminal struct {
	name   Identifier
	token  string
	source Source
}

func (t terminal) Name() Identifier {
	return t.name
}

func (t terminal) Children() []Node {
	return []Node{}
}

func (t terminal) Segment() string {
	return t.token
}

func (t terminal) Source() Source {
	return t.source
}

func (t terminal) IsTerminal() bool {
	return true
}

func NewTerminal(name Identifier, token string, cursor Cursor) Node {
	return terminal{
		name,
		token,
		Source{"", 0, cursor},
	}
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
