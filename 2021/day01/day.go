package day01

import (
	"fmt"
	"strconv"

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
	input []int
}

func New(input []string) (*Solution, error) {
	ints := []int{}
	for _, n := range input {
		num, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}
		ints = append(ints, num)
	}
	return &Solution{ints}, nil
}

func (s *Solution) Part1() (string, error) {
	increases := 0

	for i := 1; i < len(s.input); i++ {
		if s.input[i-1] < s.input[i] {
			increases++
		}
	}

	return fmt.Sprintf("%d", increases), nil
}

func (s *Solution) Part2() (string, error) {
	increases := 0

	for i := 3; i < len(s.input); i++ {
		w1 := s.input[i-3] + s.input[i-2] + s.input[i-1]
		w2 := s.input[i-2] + s.input[i-1] + s.input[i]
		if w1 < w2 {
			increases++
		}
	}

	return fmt.Sprintf("%d", increases), nil
}
