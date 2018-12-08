package main

import (
	"strconv"
	"strings"
)

type Node struct {
	Subs     []*Node
	Metadata []int
}

func (n *Node) SumMetadata() int {
	sum := 0
	for _, v := range n.Metadata {
		sum += v
	}
	for _, v := range n.Subs {
		sum += v.SumMetadata()
	}
	return sum
}

func parseData(data []int) (*Node, []int) {
	n := &Node{
		Subs: make([]*Node, 0),
	}
	nc := data[0]
	nm := data[1]
	data = data[2:]
	var cn *Node
	for i := 0; i < nc; i++ {
		cn, data = parseData(data)
		n.Subs = append(n.Subs, cn)
	}
	n.Metadata = data[:nm]
	return n, data[nm:]
}

func parseLicense(license string) *Node {
	data := make([]int, 0)
	for _, v := range strings.Split(license, " ") {
		data = append(data, mustAtoi(v))
	}
	n, _ := parseData(data)
	return n
}

func task1(in chan string) string {
	license := <-in
	tree := parseLicense(license)
	return strconv.Itoa(tree.SumMetadata())
}
