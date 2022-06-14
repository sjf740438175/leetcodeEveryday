package main

import "sort"

func main() {
	findDiagonalOrder([][]int{{2, 3}})
}

type node struct {
	row int
	col int
}

func findDiagonalOrder(mat [][]int) []int {
	m := len(mat)
	n := len(mat[0])

	var nodes []node
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			nodes = append(nodes, node{i, j})
		}
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].row+nodes[i].col < nodes[j].row+nodes[j].col || (nodes[i].row+nodes[i].col == nodes[j].row+nodes[j].col && nodes[i].row < nodes[j].row &&
			(nodes[j].row+nodes[j].col)%2 == 1) || (nodes[i].row+nodes[i].col == nodes[j].row+nodes[j].col && nodes[i].row > nodes[j].row &&
			(nodes[j].row+nodes[j].col)%2 == 0)
	})
	var ans []int
	for i := range nodes {
		ans = append(ans, mat[nodes[i].row][nodes[i].col])
	}
	return ans
}
