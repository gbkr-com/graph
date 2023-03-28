package graph

import (
	"testing"
)

type visitor struct {
	visited []int
}

func (v *visitor) visit(i int) {
	v.visited = append(v.visited, i)
}

func TestDFS(t *testing.T) {
	graph := Graph{
		0: []int{1, 2},
		1: []int{3},
		2: []int{4, 5},
	}
	v := &visitor{}
	DepthFirstSearch(graph, 0, v.visit)
	expected := []int{0, 2, 5, 4, 1, 3}
	for i, v := range v.visited {
		if v != expected[i] {
			t.Error()
		}
	}
}

func TestBFS(t *testing.T) {
	graph := Graph{
		0: []int{1, 2},
		1: []int{3},
		2: []int{4, 5},
	}
	v := &visitor{}
	BreadthFirstSearch(graph, 0, v.visit)
	expected := []int{0, 1, 2, 3, 4, 5}
	for i, v := range v.visited {
		if v != expected[i] {
			t.Error()
		}
	}
}

func TestGraph(t *testing.T) {
	// koderdojo.com
	graph := Graph{
		1: []int{2, 3},
		2: []int{4, 5},
		3: []int{5},
		4: []int{6},
		5: []int{6},
		6: []int{7},
	}
	v := &visitor{}
	DepthFirstSearch(graph, 1, v.visit)
	expected := []int{1, 3, 5, 6, 7, 2, 4}
	for i, v := range v.visited {
		if v != expected[i] {
			t.Error()
		}
	}
}
