package main

// topsort produces a listing of the vertices for all edges 'a->b', meaning 'a' comes before 'b' in the listing.
import (
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

func (g *Graph) DFS(vertex interface{}, visited map[interface{}]bool) []interface{} {
	visited[vertex] = true

	var result []interface{}
	result = append(result, vertex)

	for _, child := range g.AdjList[vertex] {
		if !visited[child] {
			out := g.DFS(child, visited)
			result = append(result, out...)
		}
	}

	return result
}

func (g *Graph) DFS2() []interface{} {
	var sorted []interface{}

	visited := make(map[interface{}]bool)
	for k := range g.AdjList {
		visited[k] = false
	}

	for vertex := range g.AdjList {
		if g.DependencyDegree(vertex) == 0 {
			out := g.DFS(vertex, visited)
			fmt.Println("Out generate by", vertex, "is", out)
			sorted = append(sorted, out)
		}
	}
	return sorted
}

func (g *Graph) TopSort() []interface{} { //Funciona mas não é o que a gnt quer, porque não gera resultados por nível só ordena as dependências

	return reverse(g.DFS2())
}

func (g *Graph) TopSort_Kahn() []interface{} {
	var sorted []interface{}
	for i := 0; i < len(g.AdjList); i++ {
		if g.DependencyDegree(g.AdjList[i]) == 0 {
			sorted = append(sorted, g.AdjList[i])
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
	fmt.Println(vertex, g.AdjList[vertex], degree)
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
	//G.AddEdge("spellchecker", "context")

	//G.AddEdge("context", "dish_syn")

	fmt.Println(G.TopSort()...)

	//go pkg
	// g := graph.New(5)
	// g.Add(2, 1)
	// fmt.Println(graph.TopSort(g))
}

// S -> D C
