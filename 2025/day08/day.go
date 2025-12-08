package day08

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/tofoss/aoc-go/pkg/registry"
	"github.com/tofoss/aoc-go/pkg/solver"
)

const year = 2025
const day = 8

func init() {
	registry.Register(year, day, func(input []string) (solver.Solver, error) {
		return New(input)
	})
}

type Point3D struct {
	X, Y, Z int
}

func NewPoint3D(x, y, z int) Point3D {
	return Point3D{
		X: x,
		Y: y,
		Z: z,
	}
}

func euclidist(a, b Point3D) float64 {
	x := (a.X - b.X)
	y := (a.Y - b.Y)
	z := (a.Z - b.Z)
	return math.Sqrt(float64(x*x + y*y + z*z))
}

type distance struct {
	dist float64
	a, b Point3D
}

func findDistances(points []Point3D) []distance {
	distances := []distance{}

	for i := range points {
		for j := i + 1; j < len(points); j++ {
			d := distance{
				dist: euclidist(points[i], points[j]),
				a:    points[i],
				b:    points[j],
			}
			distances = append(distances, d)
		}
	}
	slices.SortFunc(distances, func(a, b distance) int {
		if a.dist < b.dist {
			return -1
		} else if b.dist < a.dist {
			return 1
		} else {
			return 0
		}
	})

	return distances
}

type Solution struct {
	points []Point3D
}

func New(input []string) (*Solution, error) {
	points := make([]Point3D, len(input))
	for i := range input {
		parts := strings.Split(input[i], ",")
		x, err := strconv.Atoi(parts[0])
		y, err := strconv.Atoi(parts[1])
		z, err := strconv.Atoi(parts[2])
		if err != nil {
			return nil, err
		}
		points[i] = NewPoint3D(x, y, z)
	}
	return &Solution{points}, nil
}

func (s *Solution) Part1() (string, error) {
	distances := findDistances(s.points)
	graph := make(map[Point3D]map[Point3D]bool)
	for _, d := range distances[:10] {
		if graph[d.a] == nil {
			graph[d.a] = make(map[Point3D]bool)
		}
		if graph[d.b] == nil {
			graph[d.b] = make(map[Point3D]bool)
		}
		graph[d.a][d.b] = true
		graph[d.b][d.a] = true
	}

	circuits := findCircuits(graph)
	slices.SortFunc(circuits, func(a, b []Point3D) int {
		return len(b) - len(a)
	})

	total := 1
	for _, c := range circuits[:3] {
		total *= len(c)
	}

	return fmt.Sprintf("%d", total), nil
}

func findCircuits(graph map[Point3D]map[Point3D]bool) [][]Point3D {
	visited := make(map[Point3D]bool)
	var circuits [][]Point3D

	var dfs func(node Point3D, circuit *[]Point3D)
	dfs = func(node Point3D, circuit *[]Point3D) {
		visited[node] = true
		*circuit = append(*circuit, node)

		for neighbour := range graph[node] {
			if !visited[neighbour] {
				dfs(neighbour, circuit)
			}
		}
	}

	for node := range graph {
		if !visited[node] {
			var circuit []Point3D
			dfs(node, &circuit)
			circuits = append(circuits, circuit)
		}

	}

	return circuits
}

type UnionFind struct {
	parent map[Point3D]Point3D
	rank   map[Point3D]int
	count  int
}

func NewUnionFind(nodes []Point3D) *UnionFind {
	uf := &UnionFind{
		parent: make(map[Point3D]Point3D),
		rank:   make(map[Point3D]int),
		count:  len(nodes),
	}
	for _, node := range nodes {
		uf.parent[node] = node
		uf.rank[node] = 0
	}
	return uf
}

func (uf *UnionFind) Find(x Point3D) Point3D {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y Point3D) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return false
	}

	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}
	uf.count--
	return true
}

func findSpanningTreeFinalEdge(points []Point3D, distances []distance) distance {
	uf := NewUnionFind(points)

	for _, dist := range distances {
		conn := uf.Union(dist.a, dist.b)
		if conn && uf.count == 1 {
			return dist
		}
	}

	return distance{}
}

func (s *Solution) Part2() (string, error) {
	distances := findDistances(s.points)
	closer := findSpanningTreeFinalEdge(s.points, distances)

	return fmt.Sprintf("%d", closer.a.X*closer.b.X), nil
}
