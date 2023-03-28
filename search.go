package graph

// Graph is the map of each vertex and the vertices adjacent to it.
type Graph map[int][]int

// Visitor is a function that is called when a vertex is visited. It is never
// called more than once for a given vertex.
type Visitor func(int)

// DepthFirstSearch the given graph from the given vertex, calling the visitor
// function at each vertex.
func DepthFirstSearch(graph Graph, from int, visitor Visitor) {
	search(graph, from, visitor, pop)
}

// BreadthFirstSearch the given graph from the given vertex, calling the visitor
// function at each vertex.
func BreadthFirstSearch(graph Graph, from int, visitor Visitor) {
	search(graph, from, visitor, dequeue)
}

// Search the graph from the given vertex and call the visitor at each vertex.
// The next function returns the next vertex to visit and the remaining vertices.
// It is the only difference beween the implementations of depth first and
// breadth first searching.
func search(graph Graph, from int, visitor Visitor, next func([]int) (int, []int)) {
	visited := map[int]struct{}{}
	slice := []int{from}
	var v int
	for len(slice) > 0 {
		v, slice = next(slice)
		if _, ok := visited[v]; ok {
			continue
		}
		visited[v] = struct{}{}
		visitor(v)
		for _, e := range graph[v] {
			slice = append(slice, e)
		}
	}
}

// Pop treats the slice as a stack (LIFO).
func pop(slice []int) (int, []int) {
	n := len(slice) - 1
	return slice[n], slice[:n]
}

// Dequeue treats the slice as a queue (FIFO).
func dequeue(slice []int) (int, []int) {
	return slice[0], slice[1:]
}
