package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Node struct {
	num int
}

type Nodes []Node

func (nodes Nodes) Print() {
	for idx := range nodes {
		fmt.Println(nodes[idx].num)
	}
}

func (nodes Nodes) Len() int {
	return len(nodes)
}

func (nodes Nodes) Less(i, j int) bool {
	return nodes[i].num < nodes[j].num
}

func (nodes Nodes) Swap(i, j int) {
	nodes[i].num, nodes[j].num = nodes[j].num, nodes[i].num
}

func main() {
	nodes := Nodes{} // empty slice of nodes
	for a := 1; a < len(os.Args); a++ {
		if num, err := strconv.Atoi(os.Args[a]); err == nil {
			nodes = append(nodes, Node{num})
		}
	}

	sort.Sort(nodes)

	nodes.Print()
}
