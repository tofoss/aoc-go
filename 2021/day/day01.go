package day01

import (
	"fmt"

	"github.com/tofoss/aoc-go/pkg/readers"
	"github.com/tofoss/aoc-go/pkg/registry"
	"github.com/tofoss/aoc-go/pkg/solver"
)

const year = 2021
const day = 1

func init() {
	registry.Register(year, day, func(inputFile string) (solver.Solver, error) {
		return New(inputFile)
	})
}

type Solution struct {
	input []string
}

func New(inputFile string) (*Solution, error) {
	input, err := readers.ReadLines(inputFile)
	if err != nil {
		return nil, err
	}

	return &Solution{input}, nil
}

func (s *Solution) Part1() (string, error) {
	return fmt.Sprintf("%d-12-%02d part 1 not implemented yet\n", year, day), nil
}

func (s *Solution) Part2() (string, error) {
	return fmt.Sprintf("%d-12-%02d part 2 not implemented yet\n", year, day), nil
}
