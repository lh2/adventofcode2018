package main

import (
	"strings"
)

type Node struct {
	Name string
	Deps []*Node
	DepOf []*Node
}

func (n *Node) AddDep(dep *Node) {
	n.Deps = append(n.Deps, dep)
	dep.DepOf = append(dep.DepOf, n)
}

type Graph struct {
	Nodes map[string]*Node
}

func NewGraph() *Graph {
	var g Graph
	g.Nodes = make(map[string]*Node)
	return &g
}

func (g *Graph) Node(name string) (n *Node) {
	n, ok := g.Nodes[name]
	if ok {
		return
	}
	n = &Node{
		Name: name,
		Deps: make([]*Node, 0),
		DepOf: make([]*Node, 0),
	}
	g.Nodes[name] = n
	return
}
func (g *Graph) Remove(node *Node) {
	for _, dn := range node.DepOf {
		var i int
		var n *Node
		for i, n = range node.Deps {
			if n == node {
				break
			}
		}
		dn.Deps = append(dn.Deps[:i], dn.Deps[i+1:]...)
	}
	node.DepOf = nil
	delete(g.Nodes, node.Name)
}

func (g *Graph) Resolve() (pl []*Node) {
	pl = make([]*Node, 0, len(g.Nodes))
	nl := make([]*Node, 0, len(g.Nodes))
	for _, n := range g.Nodes {
		nl = append(nl, n)
	}
	for len(g.Nodes) > 0 {
		var ndn *Node
		for _, n := range g.Nodes {
			if len(n.Deps) == 0 && (ndn == nil || n.Name[0] < ndn.Name[0]) {
				ndn = n
			}
		}
		g.Remove(ndn)
		pl = append(pl, ndn)
	}
	return
}

func buildGraph(in chan string) (g *Graph) {
	g = NewGraph()
	for line := range in {
		lp := strings.Split(line, " ")
		n1 := g.Node(lp[1])
		n2 := g.Node(lp[7])
		n2.AddDep(n1)
	}
	return
}

func task1(in chan string) string {
	g := buildGraph(in)
	p := ""
	for _, n := range g.Resolve() {
		p += n.Name
	}
	return p
}