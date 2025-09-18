package day01

import (
	"fmt"

	"github.com/tofoss/aoc-go/pkg/registry"
	"github.com/tofoss/aoc-go/pkg/solver"
)

const year = 2021
const day = 1

func init() {
	registry.Register(year, day, func(input []string) (solver.Solver, error) {
		return New(input)
	})
}

type Solution struct {
	input []string
}

func New(input []string) (*Solution, error) {
	return &Solution{input}, nil
}

func (s *Solution) Part1() (string, error) {
	return fmt.Sprintf("%d-12-%02d part 1 not implemented yet\n", year, day), nil
}

func (s *Solution) Part2() (string, error) {
	return fmt.Sprintf("%d-12-%02d part 2 not implemented yet\n", year, day), nil
}
