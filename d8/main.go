package main

import (
	"strconv"
	"strings"
)

type Node struct {
	Subs     []*Node
	Metadata []int
}

func (n *Node) SumMetadata() (sum int) {
	for _, v := range n.Metadata {
		sum += v
	}
	for _, v := range n.Subs {
		sum += v.SumMetadata()
	}
	return
}

func (n *Node) Value() (value int) {
	if len(n.Subs) == 0 {
		for _, v := range n.Metadata {
			value += v
		}
		return
	}
	for _, v := range n.Metadata {
		i := v-1
		if i < 0 || i >= len(n.Subs) {
			continue
		}
		value += n.Subs[i].Value()
	}
	return
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

func task2(in chan string) string {
	license := <-in
	tree := parseLicense(license)
	return strconv.Itoa(tree.Value())
}