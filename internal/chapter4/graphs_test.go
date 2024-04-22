package chapter4

import (
	"fmt"
	"testing"

	utilsCh2 "github.com/fafadoboy/da-gosb/internal/chapter2/utils"
	"github.com/fafadoboy/da-gosb/internal/chapter4/models"
	"github.com/fafadoboy/da-gosb/internal/chapter4/utils"
	"github.com/samber/lo"
)

func createCityGraph() *models.Graph[models.City, float32] {
	weight := map[string]float32{"weight": 0.0}
	cityGraph := models.NewGraph[models.City, float32](models.SEATTLE, models.SANFRANCISCO, models.LOSANGELES, models.RIVERSIDE, models.PHOENIX, models.CHICAGO, models.BOSTON, models.NEWYORK, models.ATLANTA, models.MIAMI, models.DALLAS, models.HOUSTON, models.DETROIT, models.PHILADELPHIA, models.WASHINGTON)
	cityGraph.AddEdgeByVertices(models.SEATTLE, models.CHICAGO, weight)
	cityGraph.AddEdgeByVertices(models.SEATTLE, models.SANFRANCISCO, weight)
	cityGraph.AddEdgeByVertices(models.SANFRANCISCO, models.RIVERSIDE, weight)
	cityGraph.AddEdgeByVertices(models.SANFRANCISCO, models.LOSANGELES, weight)
	cityGraph.AddEdgeByVertices(models.LOSANGELES, models.RIVERSIDE, weight)
	cityGraph.AddEdgeByVertices(models.LOSANGELES, models.PHOENIX, weight)
	cityGraph.AddEdgeByVertices(models.RIVERSIDE, models.PHOENIX, weight)
	cityGraph.AddEdgeByVertices(models.RIVERSIDE, models.CHICAGO, weight)
	cityGraph.AddEdgeByVertices(models.PHOENIX, models.DALLAS, weight)
	cityGraph.AddEdgeByVertices(models.PHOENIX, models.HOUSTON, weight)
	cityGraph.AddEdgeByVertices(models.DALLAS, models.CHICAGO, weight)
	cityGraph.AddEdgeByVertices(models.DALLAS, models.ATLANTA, weight)
	cityGraph.AddEdgeByVertices(models.DALLAS, models.HOUSTON, weight)
	cityGraph.AddEdgeByVertices(models.HOUSTON, models.ATLANTA, weight)
	cityGraph.AddEdgeByVertices(models.HOUSTON, models.MIAMI, weight)
	cityGraph.AddEdgeByVertices(models.ATLANTA, models.CHICAGO, weight)
	cityGraph.AddEdgeByVertices(models.ATLANTA, models.WASHINGTON, weight)
	cityGraph.AddEdgeByVertices(models.ATLANTA, models.MIAMI, weight)
	cityGraph.AddEdgeByVertices(models.MIAMI, models.WASHINGTON, weight)
	cityGraph.AddEdgeByVertices(models.CHICAGO, models.DETROIT, weight)
	cityGraph.AddEdgeByVertices(models.DETROIT, models.BOSTON, weight)
	cityGraph.AddEdgeByVertices(models.DETROIT, models.WASHINGTON, weight)
	cityGraph.AddEdgeByVertices(models.DETROIT, models.NEWYORK, weight)
	cityGraph.AddEdgeByVertices(models.BOSTON, models.NEWYORK, weight)
	cityGraph.AddEdgeByVertices(models.NEWYORK, models.PHILADELPHIA, weight)
	cityGraph.AddEdgeByVertices(models.PHILADELPHIA, models.WASHINGTON, weight)
	return cityGraph
}

func createCityWeightedGraph() *models.Graph[models.City, float32] {
	weightedCityGraph := models.NewGraph[models.City, float32](models.SEATTLE, models.SANFRANCISCO, models.LOSANGELES, models.RIVERSIDE, models.PHOENIX, models.CHICAGO, models.BOSTON, models.NEWYORK, models.ATLANTA, models.MIAMI, models.DALLAS, models.HOUSTON, models.DETROIT, models.PHILADELPHIA, models.WASHINGTON)
	weightedCityGraph.AddEdgeByVertices(models.SEATTLE, models.CHICAGO, map[string]float32{"weight": 1737})
	weightedCityGraph.AddEdgeByVertices(models.SEATTLE, models.SANFRANCISCO, map[string]float32{"weight": 678})
	weightedCityGraph.AddEdgeByVertices(models.SANFRANCISCO, models.RIVERSIDE, map[string]float32{"weight": 386})
	weightedCityGraph.AddEdgeByVertices(models.SANFRANCISCO, models.LOSANGELES, map[string]float32{"weight": 348})
	weightedCityGraph.AddEdgeByVertices(models.LOSANGELES, models.RIVERSIDE, map[string]float32{"weight": 50})
	weightedCityGraph.AddEdgeByVertices(models.LOSANGELES, models.PHOENIX, map[string]float32{"weight": 357})
	weightedCityGraph.AddEdgeByVertices(models.RIVERSIDE, models.PHOENIX, map[string]float32{"weight": 307})
	weightedCityGraph.AddEdgeByVertices(models.RIVERSIDE, models.CHICAGO, map[string]float32{"weight": 1704})
	weightedCityGraph.AddEdgeByVertices(models.PHOENIX, models.DALLAS, map[string]float32{"weight": 887})
	weightedCityGraph.AddEdgeByVertices(models.PHOENIX, models.HOUSTON, map[string]float32{"weight": 1015})
	weightedCityGraph.AddEdgeByVertices(models.DALLAS, models.CHICAGO, map[string]float32{"weight": 805})
	weightedCityGraph.AddEdgeByVertices(models.DALLAS, models.ATLANTA, map[string]float32{"weight": 721})
	weightedCityGraph.AddEdgeByVertices(models.DALLAS, models.HOUSTON, map[string]float32{"weight": 225})
	weightedCityGraph.AddEdgeByVertices(models.HOUSTON, models.ATLANTA, map[string]float32{"weight": 702})
	weightedCityGraph.AddEdgeByVertices(models.HOUSTON, models.MIAMI, map[string]float32{"weight": 968})
	weightedCityGraph.AddEdgeByVertices(models.ATLANTA, models.CHICAGO, map[string]float32{"weight": 588})
	weightedCityGraph.AddEdgeByVertices(models.ATLANTA, models.WASHINGTON, map[string]float32{"weight": 543})
	weightedCityGraph.AddEdgeByVertices(models.ATLANTA, models.MIAMI, map[string]float32{"weight": 604})
	weightedCityGraph.AddEdgeByVertices(models.MIAMI, models.WASHINGTON, map[string]float32{"weight": 923})
	weightedCityGraph.AddEdgeByVertices(models.CHICAGO, models.DETROIT, map[string]float32{"weight": 238})
	weightedCityGraph.AddEdgeByVertices(models.DETROIT, models.BOSTON, map[string]float32{"weight": 613})
	weightedCityGraph.AddEdgeByVertices(models.DETROIT, models.WASHINGTON, map[string]float32{"weight": 396})
	weightedCityGraph.AddEdgeByVertices(models.DETROIT, models.NEWYORK, map[string]float32{"weight": 482})
	weightedCityGraph.AddEdgeByVertices(models.BOSTON, models.NEWYORK, map[string]float32{"weight": 190})
	weightedCityGraph.AddEdgeByVertices(models.NEWYORK, models.PHILADELPHIA, map[string]float32{"weight": 81})
	weightedCityGraph.AddEdgeByVertices(models.PHILADELPHIA, models.WASHINGTON, map[string]float32{"weight": 123})
	return weightedCityGraph
}

func converter(f float32) float32 { return f }

func TestGraphCreation(t *testing.T) {
	cityGraph := createCityGraph()
	fmt.Println(cityGraph.ToString())
}

func TestGraphsNaiveShortestPath(t *testing.T) {
	cityGraph := createCityGraph()
	fmt.Println(cityGraph.ToString())

	algo := utilsCh2.AlgoSearch[models.City]{}
	if sol := algo.BFS(models.BOSTON, func(c models.City) bool { return c == models.MIAMI }, cityGraph.NeighborsForVertex); sol != nil {
		path := sol.ToPath()
		fmt.Printf("Path from %s to %s:\n%v", models.BOSTON, models.MIAMI, path)
	} else {
		fmt.Println("No solution found using depth-first rearch")
	}
}

func TestGraphsShortestPath(t *testing.T) {
	cityGraph := createCityWeightedGraph()
	fmt.Println(cityGraph.ToString())

	algo := utils.AlgoGraph[models.City, float32]{}
	if sol := algo.MST(cityGraph, 0, converter); sol != nil {
		var sum float32

		for _, edge := range sol {
			weight := converter(edge.Meta["weight"])
			fmt.Printf("%v (%v) > %v\n", cityGraph.VertexAt(edge.U), weight, cityGraph.VertexAt(edge.V))
			sum += weight
		}
		fmt.Printf("Total Weight: %v\n", sum)
		// fmt.Printf("Path from %s to %s:\n%v", models.BOSTON, models.MIAMI, path)
	} else {
		fmt.Println("No solution found using MST")
	}
}

func TestGraphsShortestPathWithDjakstra(t *testing.T) {
	cityGraph := createCityWeightedGraph()
	fmt.Println(cityGraph.ToString())

	algo := utils.AlgoGraph[models.City, float32]{}
	distances, path := algo.Dijkstra(cityGraph, models.LOSANGELES, converter)

	// display the results
	lo.ForEach(distances, func(item *float32, index int) {
		fmt.Printf("%s : %v\n", cityGraph.VertexAt(index), *item)
	})
	fmt.Println()

	start := cityGraph.IndexOf(models.LOSANGELES)
	end := cityGraph.IndexOf(models.BOSTON)
	edgesPath := make([]*models.Edge[float32], 0)

	e := path[end]
	edgesPath = append(edgesPath, e)
	for e.U != start {
		e = path[e.U]
		edgesPath = append(edgesPath, e)
	}
	var sum float32
	for _, edge := range lo.Reverse[*models.Edge[float32]](edgesPath) {
		weight := converter(edge.Meta["weight"])
		fmt.Printf("%v %+v> %v\n", cityGraph.VertexAt(edge.U), weight, cityGraph.VertexAt(edge.V))
		sum += weight
	}
	fmt.Printf("Total Weight: %v\n", sum)
}
