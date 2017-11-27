package pov

type children []string

type Graph struct {
	root  string
	nodes []string
	arcs  map[string]children
}

func New() *Graph {
	return &Graph{"parent", []string{"parent"}, map[string]children{}}
}

func (graph *Graph) AddNode(nodeLabel string) {
	graph.nodes = append(graph.nodes, nodeLabel)
}

func (graph *Graph) AddArc(from, to string) {
	graph.arcs[from] = append(graph.arcs[from], to)
}

func (graph *Graph) ArcList() []string {
	result := []string{}
	for from, nodeArcs := range graph.arcs {
		for _, to := range nodeArcs {
			result = append(result, from+" -> "+to)
		}
	}
	return result
}

func (graph *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	newGraph := graph.copy()
	if path, found := newGraph.find(oldRoot, newRoot); found {
		for i := len(path) - 1; i > 0; i-- {
			from, to := path[i], path[i-1]
			newGraph.arcs[from] = newGraph.arcs[from].remove(to)
			newGraph.arcs[to] = append(newGraph.arcs[to], path[i])
		}
	}
	return newGraph
}

func (graph *Graph) find(root, toFind string) (path []string, found bool) {
	for _, node := range graph.arcs[root] {
		if node == toFind {
			path = append(path, node, root)
			found = true
			return
		}
	}
	for _, node := range graph.arcs[root] {
		path, found = graph.find(node, toFind)
		if found {
			path = append(path, root)
			return
		}
	}
	return nil, false
}

func (graph *Graph) copy() *Graph {
	newGraph := Graph{}
	newGraph.root = graph.root

	newGraph.nodes = make([]string, len(graph.nodes))
	copy(newGraph.nodes, graph.nodes)

	newGraph.arcs = make(map[string]children)
	for from, toNodes := range graph.arcs {
		newGraph.arcs[from] = make([]string, len(toNodes))
		copy(newGraph.arcs[from], toNodes)
	}
	return &newGraph
}

func (toNodes children) remove(node string) children {
	for i, child := range toNodes {
		if child == node {
			return append(toNodes[:i], toNodes[i+1:]...)
		}
	}
	return toNodes
}
