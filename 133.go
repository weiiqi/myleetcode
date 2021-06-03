package main

type Nodex struct {
	Val       int
	Neighbors []*Nodex
}

func cloneGraph(node *Nodex) *Nodex {
	visited := make(map[*Nodex]*Nodex)
	return clone(node, visited)
}

// 思路L：用个map存储节点是否已经被复制。递归复制节点。
func clone(node *Nodex, visited map[*Nodex]*Nodex) *Nodex {
	if node == nil {
		return nil
	}
	if v, ok := visited[node]; ok {
		return v
	}
	newNode := &Nodex{
		Val:       node.Val,
		Neighbors: make([]*Nodex, len(node.Neighbors)),
	}
	visited[node] = newNode
	for i := 0; i < len(node.Neighbors); i++ {
		newNode.Neighbors[i] = clone(node.Neighbors[i], visited)
	}
	return newNode
}
