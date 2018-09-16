package v2

//
// Generating output using ASTs
//

type Output interface{}

type NodeGenerator func(Node, []Output) Output

func GenerateAll(root Node, generator NodeGenerator) Output {
	return nil
}

func FoldUp(root Node, generators []struct {
	NodePredicate
	NodeGenerator
}) Output {
	return nil
}

func IsA(name string) NodePredicate {
	// IsA("FORLOOP")(node) -> T/F
	return func(n Node) bool {
		return n.Name() == name
	}
}

func Map(nodes []Node, generator NodeGenerator) []Output {
	out := make([]Output, len(nodes))
	for i, node := range nodes {
		out[i] = generator(node, nil)
	}
	return out
}
