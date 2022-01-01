package graph

type Graph struct {
	Nodes map[string]*Node
}

type Node struct {
	id    string
	conTo []*Node
}

func NewNode(id string) *Node {
	return &Node{
		id:    id,
		conTo: []*Node{},
	}
}

func (g *Graph) AddNode(id string) {
	if _, ok := g.Nodes[id]; !ok {
		n := NewNode(id)
		g.Nodes[id] = n
	}

}

func (g *Graph) AddEdge(srcId, destId string, directed bool) {
	g.AddNode(srcId)
	g.AddNode(destId)
	srcNode := g.GetNode(srcId)
	destNode := g.GetNode(destId)
	if srcNode.isConnectedTo(destNode) {
		return
	}

	g.Nodes[srcId].conTo = append(g.Nodes[srcId].conTo, g.Nodes[destId])
	if !directed {
		if destNode.isConnectedTo(srcNode) {
			return
		}
		destNode.conTo = append(destNode.conTo, srcNode)
	}
}

func (n *Node) isConnectedTo(destNode *Node) bool {
	for _, node := range n.conTo {
		if node.id == destNode.id {
			return true
		}
	}
	return false

}

func (g *Graph) GetNode(id string) *Node {
	return g.Nodes[id]
}

func (n *Node) GetConnections() []*Node {
	return n.conTo
}

func (n *Node) GetId() string {
	return n.id
}
