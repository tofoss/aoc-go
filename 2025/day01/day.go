package day01

import (
	"fmt"
	"strconv"

	"github.com/tofoss/aoc-go/pkg/registry"
	"github.com/tofoss/aoc-go/pkg/solver"
)

const year = 2025
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
	dial := 50
	count := 0
	for _, in := range s.input {
		direction := in[0]
		steps, err := strconv.Atoi(in[1:])
		if err != nil {
			return "", err
		}
		dial = rotate(dial, steps, direction)
		if dial == 0 {
			count++
		}
	}
	return fmt.Sprintf("%d", count), nil
}

func (s *Solution) Part2() (string, error) {
	dial := 50
	count := 0
	for _, in := range s.input {
		direction := in[0]
		steps, err := strconv.Atoi(in[1:])
		if steps > 100 {
			count += steps / 100
			steps = steps % 100
		}
		if err != nil {
			return "", err
		}
		p0 := false
		dial, p0 = rotate2(dial, steps, direction)
		if dial == 0 {
			count++
		}
		if p0 {
			count++
		}
	}
	return fmt.Sprintf("%d", count), nil
}

func rotate(dial, steps int, direction byte) int {
	if direction == 'R' {
		dial += steps
	} else {
		dial -= steps
		if dial < 0 {
			dial = 100 + dial
		}
	}
	return dial % 100
}

func rotate2(dial, steps int, direction byte) (int, bool) {
	p0 := false
	if direction == 'R' {
		next_dial := dial + steps
		if dial != 100 && next_dial > 100 {
			p0 = true
		}
		return next_dial % 100, p0
	} else {
		next_dial := dial - steps
		if next_dial < 0 {
			next_dial = 100 + next_dial
			if dial != 0 {
				p0 = true
			}
		}
		return next_dial % 100, p0
	}
}
