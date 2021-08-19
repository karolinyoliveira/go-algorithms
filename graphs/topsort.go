package main

// topsort produces a listing of the vertices for all edges 'a->b', meaning 'a' comes before 'b' in the listing.
import (
	"container/list"
	"fmt"
)

const (
	white = iota
	grey
	black
)

type Graph struct {
	AdjList map[interface{}][]interface{}
}

func InitGraph() (g *Graph) {
	return &Graph{
		AdjList: map[interface{}][]interface{}{},
	}
}

func (g *Graph) AddVertex(src interface{}) {
	g.AdjList[src] = make([]interface{}, 0)
}

func (g *Graph) AddEdge(src interface{}, dst interface{}) {
	if _, ok := g.AdjList[src]; ok {
		g.AdjList[src] = append(g.AdjList[src], dst)
	} else {
		fmt.Println("Source vertex does not exists")
	}
}

func (g *Graph) DFSRec(vertex interface{}, visited map[interface{}]bool) []interface{} {
	visited[vertex] = true

	var result []interface{}
	result = append(result, vertex)

	for _, child := range g.AdjList[vertex] {
		if !visited[child] {
			out := g.DFSRec(child, visited)
			result = append(result, out...)
		}
	}

	return result
}

func (g *Graph) DFS() []interface{} {
	var sorted []interface{}

	visited := make(map[interface{}]bool)
	for k := range g.AdjList {
		visited[k] = false
	}

	for vertex := range g.AdjList {
		if g.DependencyDegree(vertex) == 0 {
			out := g.DFSRec(vertex, visited)
			fmt.Println("Out generate by", vertex, "is", out)
			sorted = append(sorted, out)
		}
	}
	return sorted
}

func (g *Graph) TopSort() []interface{} {
	return reverse(g.DFS())
}

func (g *Graph) TopSort_Kahn() []interface{} { //By degree
	var sorted []interface{}
	degrees := make(map[interface{}]int)
	queue := list.New()
	//Get degree of every vertex
	for v := range g.AdjList {
		degrees[v] = g.DependencyDegree(v)
	}

	for v := range g.AdjList {
		if degrees[v] == 0 {
			queue.PushBack(v)

		}
	}
	fmt.Println("queue:", queue)
	fmt.Println("degrees:", degrees)
	for queue.Len() > 0 {
		currVertex := queue.Front()
		queue.Remove(currVertex)
		sorted = append(sorted, currVertex.Value)

		for _, child := range g.AdjList[currVertex.Value] {
			fmt.Println("childs:", g.AdjList[currVertex.Value], "of", currVertex.Value)
			degrees[child]--

			if degrees[child] == 0 {
				queue.PushBack(child)
			}
		}
	}

	return sorted
}

func reverse(vec []interface{}) []interface{} {
	reversed := make([]interface{}, len(vec))
	copy(reversed, vec)

	for i := len(reversed)/2 - 1; i >= 0; i-- {
		j := len(reversed) - 1 - i
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}

	return reversed
}

func remove(el interface{}, slice []interface{}) []interface{} {
	for i, v := range slice {
		if el == v {
			slice[i] = slice[len(slice)-1]
			return slice[:len(slice)-1]
		}
	}

	aux := slice[0]
	if len(slice) == 2 {
		slice = make([]interface{}, 0)
		slice = append(slice, aux)
	}
	fmt.Println(len(slice))
	return slice
}

func (g *Graph) EdgeExists(src interface{}, dst interface{}) bool {
	for _, v := range g.AdjList[src] {
		if v == dst {
			return true
		}
	}
	return false
}

func (g *Graph) DependencyDegree(vertex interface{}) int {
	degree := 0

	for v := range g.AdjList {
		if g.EdgeExists(v, vertex) {
			degree++
		}
	}
	return degree
}

func main() {
	G := InitGraph()

	models := []interface{}{
		"spellchecker",
		"dish_syn",
		"context",
	}

	for i := 0; i < len(models); i++ {
		G.AddVertex(models[i])
	}

	G.AddEdge("spellchecker", "dish_syn")

	fmt.Println(G.TopSort_Kahn()) // [context spellchecker dish_syn] ou [spellchecker context dish_syn]
}
