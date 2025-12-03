package day03

import (
	"strconv"

	"github.com/tofoss/aoc-go/pkg/registry"
	"github.com/tofoss/aoc-go/pkg/slice"
	"github.com/tofoss/aoc-go/pkg/solver"
)

const year = 2025
const day = 3

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
	return s.solve(2, s.keepTopTwo)
}

func (s *Solution) Part2() (string, error) {
	return s.solve(12, s.removeDescending)
}

func (s *Solution) solve(bufferSize int, updateBuffer func([]rune, rune)) (string, error) {
	buffer := make([]rune, bufferSize)
	results := make([]string, 0, len(s.input))

	for _, row := range s.input {
		clear(buffer)
		for _, col := range row {
			updateBuffer(buffer, col)
		}
		results = append(results, string(buffer))
	}

	sum, err := slice.SumOfStrs(results)
	return strconv.Itoa(sum), err
}

func (s *Solution) keepTopTwo(buffer []rune, next rune) {
	if buffer[1] > buffer[0] {
		buffer[0] = buffer[1]
		buffer[1] = 0
	}
	if next > buffer[1] {
		buffer[1] = next
	}
}

func (s *Solution) removeDescending(buffer []rune, next rune) {
	for i := 0; i < len(buffer)-1; i++ {
		if buffer[i] < buffer[i+1] {
			copy(buffer[i:], buffer[i+1:])
			buffer[len(buffer)-1] = 0
			break
		}
	}
	if buffer[len(buffer)-1] < next {
		buffer[len(buffer)-1] = next
	}
}
