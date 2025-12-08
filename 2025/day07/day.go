package day07

import (
	"fmt"

	"github.com/tofoss/aoc-go/pkg/grid"
	"github.com/tofoss/aoc-go/pkg/registry"
	"github.com/tofoss/aoc-go/pkg/solver"
)

const year = 2025
const day = 7

func init() {
	registry.Register(year, day, func(input []string) (solver.Solver, error) {
		return New(input)
	})
}

type Solution struct {
	matrix [][]rune
}

type Node struct {
	Edges []Node
	Value grid.Point
}

func New(input []string) (*Solution, error) {
	matrix := make([][]rune, len(input))
	for i := range input {
		matrix[i] = []rune(input[i])
	}
	return &Solution{matrix}, nil
}

func (s *Solution) Part1() (string, error) {
	start := grid.Point{}
	for i := range s.matrix[0] {
		if s.matrix[0][i] == 'S' {
			start = grid.Point{Y: 0, X: i}
		}
	}
	movesMade := make(map[grid.Point]int)
	nextMoves := make(map[grid.Point]int)
	nextMoves[start] = 1
	splitCount := 0

	for range s.matrix {
		nextNext := make(map[grid.Point]int)
		for move := range nextMoves {
			down := move.Down()
			movesMade[move] = 1
			if grid.OutOfBounds(down, s.matrix) {
				continue
			} else if s.matrix[down.Y][down.X] == '.' {
				nextNext[down] = 1
			} else if s.matrix[down.Y][down.X] == '^' {
				splitCount++
				nextNext[move.DownLeft()] = 1
				nextNext[move.DownRight()] = 1
			}
		}
		nextMoves = nextNext
	}

	//for move := range movesMade {
	//	s.matrix[move.Y][move.X] = '|'
	//}
	//for _, r := range s.matrix {
	//	fmt.Println(string(r))
	//}

	return fmt.Sprintf("%d", splitCount), nil
}

func (s *Solution) Part2() (string, error) {
	start := grid.Point{}
	for i := range s.matrix[0] {
		if s.matrix[0][i] == 'S' {
			start = grid.Point{Y: 0, X: i}
		}
	}
	movesMade := make(map[grid.Point]int)
	nextMoves := make(map[grid.Point]int)
	nextMoves[start] = 1
	graph := make(map[grid.Point][]grid.Point)

	for range s.matrix {
		nextNext := make(map[grid.Point]int)
		for move := range nextMoves {
			down := move.Down()
			movesMade[move] = 1
			if grid.OutOfBounds(down, s.matrix) {
				continue
			} else if s.matrix[down.Y][down.X] == '.' {
				nextNext[down] = 1
				graph[move] = append(graph[move], down)
			} else if s.matrix[down.Y][down.X] == '^' {
				nextNext[move.DownLeft()] = 1
				nextNext[move.DownRight()] = 1
				graph[move] = append(graph[move], move.DownLeft())
				graph[move] = append(graph[move], move.DownRight())
			}
		}
		nextMoves = nextNext
	}

	paths := findAllPaths(graph, start)

	return fmt.Sprintf("%d", len(paths)), nil
}

func findAllPaths(
	graph map[grid.Point][]grid.Point,
	start grid.Point,
) [][]grid.Point {
	var paths [][]grid.Point
	var currentPath []grid.Point
	visited := make(map[grid.Point]bool)

	var dfs func(node grid.Point)
	dfs = func(node grid.Point) {
		currentPath = append(currentPath, node)
		visited[node] = true

		if len(graph[node]) == 0 {
			path := make([]grid.Point, len(currentPath))
			copy(path, currentPath)
			paths = append(paths, path)
		} else {
			for _, neighbour := range graph[node] {
				if !visited[neighbour] {
					dfs(neighbour)
				}
			}
		}
		currentPath = currentPath[:len(currentPath)-1]
		visited[node] = false
	}

	dfs(start)
	return paths
}
