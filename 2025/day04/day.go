package day04

import (
	"fmt"

	"github.com/tofoss/aoc-go/pkg/grid"
	"github.com/tofoss/aoc-go/pkg/registry"
	"github.com/tofoss/aoc-go/pkg/solver"
)

const year = 2025
const day = 4

func init() {
	registry.Register(year, day, func(input []string) (solver.Solver, error) {
		return New(input)
	})
}

type Solution struct {
	input [][]rune
}

func New(input []string) (*Solution, error) {
	grid := make([][]rune, len(input))
	for i, line := range input {
		grid[i] = []rune(line)
	}
	return &Solution{grid}, nil
}

func accessibleRolls(g [][]rune) ([][]rune, int) {
	gridRows := len(g) - 1
	gridCols := len(g[0]) - 1
	accessibleRolls := 0
	newGrid := make([][]rune, len(g))
	for y := range g {
		for x := range g[y] {
			adjacentRolls := 0
			for _, adj := range grid.Adjencencies {
				point := grid.Point{Y: adj.Y + y, X: adj.X + x}
				if !point.OffGrid(gridRows, gridCols) && g[point.Y][point.X] == '@' {
					adjacentRolls++
				}
			}
			if g[y][x] == '@' && adjacentRolls < 4 {
				accessibleRolls++
				newGrid[y] = append(newGrid[y], 'x')
			} else {
				newGrid[y] = append(newGrid[y], g[y][x])
			}
		}
	}

	return newGrid, accessibleRolls
}

func (s *Solution) Part1() (string, error) {
	_, rolls := accessibleRolls(s.input)
	return fmt.Sprintf("%d", rolls), nil
}

func (s *Solution) Part2() (string, error) {
	grid := s.input
	sum := 0
	for {
		rolls := 0
		grid, rolls = accessibleRolls(grid)
		if rolls == 0 {
			break
		}
		sum += rolls
	}

	return fmt.Sprintf("%d", sum), nil
}
