package v2

type Source struct {
	File  string
	Line  int
	Index int
}

type Node interface {
	Name() string
	Children() []Node
	Raw() string
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

func (t *terminal) Raw() string {
	return t.token
}

func (t *terminal) Source() Source {
	return t.source
}

func (t *terminal) IsTerminal() bool {
	return true
}

func NewTerminal(name string, token string, source Source) Node {
	return &terminal{
		name,
		token,
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
