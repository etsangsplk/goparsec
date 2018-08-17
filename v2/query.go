package v2

func QueryFilter(root Node, predicate NodePredicate) []Node {
	return nil
}

func QueryTerminals(root Node) []Node {
	return QueryFilter(root, func(n Node) bool {
		return n.IsTerminal()
	})
}

type QueryString string

func QueryOne(root Node, query QueryString) Node {
	return nil
}

func QueryMany(root Node, query QueryString) []Node {
	return nil
}

func Walk(root Node) chan Node {
	return nil
}
