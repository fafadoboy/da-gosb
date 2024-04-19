package chapter4

import (
	"fmt"
	"testing"

	"github.com/fafadoboy/da-gosb/internal/chapter2/utils"
	"github.com/fafadoboy/da-gosb/internal/chapter4/models"
)

func createCityGraph() *models.Graph[models.City] {
	cityGraph := models.NewGraph[models.City](models.SEATTLE, models.SANFRANCISCO, models.LOSANGELES, models.RIVERSIDE, models.PHOENIX, models.CHICAGO, models.BOSTON, models.NEWYORK, models.ATLANTA, models.MIAMI, models.DALLAS, models.HOUSTON, models.DETROIT, models.PHILADELPHIA, models.WASHINGTON)
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
func TestGraphCreation(t *testing.T) {
	// city_graph: Graph[str] = Graph([models.SEATTLE, models.SANFRANCISCO, "LosAngeles", models.RIVERSIDE, models.PHOENIX, models.CHICAGO, models.BOSTON, models.NEWYORK, models.ATLANTA, models.MIAMI, models.DALLAS, models.HOUSTON, models.DETROIT, models.PHILADELPHIA, models.WASHINGTON])
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
