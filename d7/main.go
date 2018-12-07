package main

import (
	"strconv"
	"strings"
)

type Node struct {
	Name  string
	Deps  []*Node
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
		Name:  name,
		Deps:  make([]*Node, 0),
		DepOf: make([]*Node, 0),
	}
	g.Nodes[name] = n
	return
}

func nodeSliceIndexOf(nodes []*Node, node *Node) int {
	i := -1
	var n *Node
	for i, n = range nodes {
		if n == node {
			break
		}
	}
	return i
}

func (g *Graph) Remove(node *Node) {
	for _, dn := range node.DepOf {
		i := nodeSliceIndexOf(dn.Deps, node)
		dn.Deps = append(dn.Deps[:i], dn.Deps[i+1:]...)
	}
	node.DepOf = nil
	delete(g.Nodes, node.Name)
}

func (g *Graph) NoDepNodes() (ns []*Node) {
	ns = make([]*Node, 0)
	for _, n := range g.Nodes {
		if len(n.Deps) == 0 {
			ns = append(ns, n)
		}
	}
	return ns
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

type Worker struct {
	Task          *Node
	TimeRemaining int
}

type Workers []*Worker

func NewWorkers() Workers {
	return []*Worker{
		&Worker{},
		&Worker{},
		&Worker{},
		&Worker{},
		&Worker{},
	}
}

func (ws Workers) FreeWorker() *Worker {
	var fw *Worker
	for _, w := range ws {
		if w.TimeRemaining == 0 {
			fw = w
			break
		}
	}
	return fw
}

func (ws Workers) Tick() {
	for _, w := range ws {
		if w.TimeRemaining > 0 {
			w.TimeRemaining--
		}
	}
}

func (ws Workers) NotAssigned(tasks []*Node) []*Node {
	for _, w := range ws {
		if w.Task == nil {
			continue
		}
		i := nodeSliceIndexOf(tasks, w.Task)
		if i < 0 {
			continue
		}
		tasks = append(tasks[:i], tasks[i+1:]...)
	}
	return tasks
}

func task2(in chan string) string {
	g := buildGraph(in)
	ws := NewWorkers()

	i := -1
	tasks := g.NoDepNodes()
	for ; len(g.Nodes) > 0; i++ {
		ws.Tick()
		taskFinished := false
		for _, w := range ws {
			if w.Task != nil && w.TimeRemaining == 0 {
				g.Remove(w.Task)
				w.Task = nil
				taskFinished = true
			}
		}
		if taskFinished {
			tasks = g.NoDepNodes()
			tasks = ws.NotAssigned(tasks)
		}

		taskAssigned := false
		for _, task := range tasks {
			w := ws.FreeWorker()
			if w == nil {
				break
			}
			w.Task = task
			w.TimeRemaining = 60 + int(task.Name[0]-'A'+1)
			taskAssigned = true
		}
		if taskAssigned {
			tasks = ws.NotAssigned(tasks)
		}
	}
	return strconv.Itoa(i)
}
