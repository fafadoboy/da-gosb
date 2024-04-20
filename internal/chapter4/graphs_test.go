package chapter4

import (
	"fmt"
	"testing"

	"github.com/fafadoboy/da-gosb/internal/chapter2/utils"
	"github.com/fafadoboy/da-gosb/internal/chapter4/models"
)

func createCityGraph() *models.Graph[models.City] {
	cityGraph := models.NewGraph(models.SEATTLE, models.SANFRANCISCO, models.LOSANGELES, models.RIVERSIDE, models.PHOENIX, models.CHICAGO, models.BOSTON, models.NEWYORK, models.ATLANTA, models.MIAMI, models.DALLAS, models.HOUSTON, models.DETROIT, models.PHILADELPHIA, models.WASHINGTON)
	cityGraph.AddEdgeByVertices(models.SEATTLE, models.CHICAGO)
	cityGraph.AddEdgeByVertices(models.SEATTLE, models.SANFRANCISCO)
	cityGraph.AddEdgeByVertices(models.SANFRANCISCO, models.RIVERSIDE)
	cityGraph.AddEdgeByVertices(models.SANFRANCISCO, models.LOSANGELES)
	cityGraph.AddEdgeByVertices(models.LOSANGELES, models.RIVERSIDE)
	cityGraph.AddEdgeByVertices(models.LOSANGELES, models.PHOENIX)
	cityGraph.AddEdgeByVertices(models.RIVERSIDE, models.PHOENIX)
	cityGraph.AddEdgeByVertices(models.RIVERSIDE, models.CHICAGO)
	cityGraph.AddEdgeByVertices(models.PHOENIX, models.DALLAS)
	cityGraph.AddEdgeByVertices(models.PHOENIX, models.HOUSTON)
	cityGraph.AddEdgeByVertices(models.DALLAS, models.CHICAGO)
	cityGraph.AddEdgeByVertices(models.DALLAS, models.ATLANTA)
	cityGraph.AddEdgeByVertices(models.DALLAS, models.HOUSTON)
	cityGraph.AddEdgeByVertices(models.HOUSTON, models.ATLANTA)
	cityGraph.AddEdgeByVertices(models.HOUSTON, models.MIAMI)
	cityGraph.AddEdgeByVertices(models.ATLANTA, models.CHICAGO)
	cityGraph.AddEdgeByVertices(models.ATLANTA, models.WASHINGTON)
	cityGraph.AddEdgeByVertices(models.ATLANTA, models.MIAMI)
	cityGraph.AddEdgeByVertices(models.MIAMI, models.WASHINGTON)
	cityGraph.AddEdgeByVertices(models.CHICAGO, models.DETROIT)
	cityGraph.AddEdgeByVertices(models.DETROIT, models.BOSTON)
	cityGraph.AddEdgeByVertices(models.DETROIT, models.WASHINGTON)
	cityGraph.AddEdgeByVertices(models.DETROIT, models.NEWYORK)
	cityGraph.AddEdgeByVertices(models.BOSTON, models.NEWYORK)
	cityGraph.AddEdgeByVertices(models.NEWYORK, models.PHILADELPHIA)
	cityGraph.AddEdgeByVertices(models.PHILADELPHIA, models.WASHINGTON)
	return cityGraph
}

func createCityWeightedGraph() *models.WeightedGraph[models.City] {
	weightedCityGraph := models.NewWightedGraph(models.SEATTLE, models.SANFRANCISCO, models.LOSANGELES, models.RIVERSIDE, models.PHOENIX, models.CHICAGO, models.BOSTON, models.NEWYORK, models.ATLANTA, models.MIAMI, models.DALLAS, models.HOUSTON, models.DETROIT, models.PHILADELPHIA, models.WASHINGTON)
	weightedCityGraph.AddEdgeByVertices(models.SEATTLE, models.CHICAGO, 1737)
	weightedCityGraph.AddEdgeByVertices(models.SEATTLE, models.SANFRANCISCO, 678)
	weightedCityGraph.AddEdgeByVertices(models.SANFRANCISCO, models.RIVERSIDE, 386)
	weightedCityGraph.AddEdgeByVertices(models.SANFRANCISCO, models.LOSANGELES, 348)
	weightedCityGraph.AddEdgeByVertices(models.LOSANGELES, models.RIVERSIDE, 50)
	weightedCityGraph.AddEdgeByVertices(models.LOSANGELES, models.PHOENIX, 357)
	weightedCityGraph.AddEdgeByVertices(models.RIVERSIDE, models.PHOENIX, 307)
	weightedCityGraph.AddEdgeByVertices(models.RIVERSIDE, models.CHICAGO, 1704)
	weightedCityGraph.AddEdgeByVertices(models.PHOENIX, models.DALLAS, 887)
	weightedCityGraph.AddEdgeByVertices(models.PHOENIX, models.HOUSTON, 1015)
	weightedCityGraph.AddEdgeByVertices(models.DALLAS, models.CHICAGO, 805)
	weightedCityGraph.AddEdgeByVertices(models.DALLAS, models.ATLANTA, 721)
	weightedCityGraph.AddEdgeByVertices(models.DALLAS, models.HOUSTON, 225)
	weightedCityGraph.AddEdgeByVertices(models.HOUSTON, models.ATLANTA, 702)
	weightedCityGraph.AddEdgeByVertices(models.HOUSTON, models.MIAMI, 968)
	weightedCityGraph.AddEdgeByVertices(models.ATLANTA, models.CHICAGO, 588)
	weightedCityGraph.AddEdgeByVertices(models.ATLANTA, models.WASHINGTON, 543)
	weightedCityGraph.AddEdgeByVertices(models.ATLANTA, models.MIAMI, 604)
	weightedCityGraph.AddEdgeByVertices(models.MIAMI, models.WASHINGTON, 923)
	weightedCityGraph.AddEdgeByVertices(models.CHICAGO, models.DETROIT, 238)
	weightedCityGraph.AddEdgeByVertices(models.DETROIT, models.BOSTON, 613)
	weightedCityGraph.AddEdgeByVertices(models.DETROIT, models.WASHINGTON, 396)
	weightedCityGraph.AddEdgeByVertices(models.DETROIT, models.NEWYORK, 482)
	weightedCityGraph.AddEdgeByVertices(models.BOSTON, models.NEWYORK, 190)
	weightedCityGraph.AddEdgeByVertices(models.NEWYORK, models.PHILADELPHIA, 81)
	weightedCityGraph.AddEdgeByVertices(models.PHILADELPHIA, models.WASHINGTON, 123)
	return weightedCityGraph
}

func TestGraphCreation(t *testing.T) {
	cityGraph := createCityGraph()
	fmt.Println(cityGraph.ToString())
}

func TestGraphsNaiveShortestPath(t *testing.T) {
	cityGraph := createCityGraph()
	fmt.Println(cityGraph.ToString())

	algo := utils.AlgoSearch[models.City]{}
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
}
