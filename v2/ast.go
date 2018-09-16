package v2

type Source struct {
	File   string
	Line   int
	Index  int
	Cursor int
}

type Node interface {
	Name() string
	Children() []Node
	Value() string
	Source() Source
	IsTerminal() bool
}

//
// Terminal
//
type terminal struct {
	name   string
	token  string
	source Source
}

func (t *terminal) Name() string {
	return t.name
}

func (t *terminal) Children() []Node {
	return []Node{}
}

func (t *terminal) Value() string {
	return t.token
}

func (t *terminal) Source() Source {
	return t.source
}

func (t *terminal) IsTerminal() bool {
	return true
}

// NewTerminal will create and return a new terminal node.
func NewTerminal(name string, value string, source Source) Node {
	return &terminal{
		name,
		value,
		source,
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
