package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tofoss/aoc-go/pkg/math"
	"github.com/tofoss/aoc-go/pkg/registry"
	"github.com/tofoss/aoc-go/pkg/solver"
)

const year = 2025
const day = 2

func init() {
	registry.Register(year, day, func(input []string) (solver.Solver, error) {
		return New(input)
	})
}

type Solution struct {
	ranges []Range
}

type Range struct {
	From, To int
}

func New(input []string) (*Solution, error) {
	inputRanges := strings.Split(input[0], ",")
	ranges := []Range{}
	for _, r := range inputRanges {
		parts := strings.Split(r, "-")
		from, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}
		to, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}

		ranges = append(ranges, Range{from, to})
	}

	return &Solution{ranges}, nil
}

func (s *Solution) Part1() (string, error) {
	sum := s.sumIf(func(i int) bool {
		str := strconv.Itoa(i)
		if len(str)%2 != 0 {
			return false
		}

		first := str[:len(str)/2]
		second := str[len(str)/2:]
		return first == second
	})
	return fmt.Sprintf("%d", sum), nil
}

func (s *Solution) Part2() (string, error) {
	sum := s.sumIf(func(i int) bool {
		return isRepeating(strconv.Itoa(i))
	})
	return fmt.Sprintf("%d", sum), nil
}

func (s *Solution) sumIf(predicate func(int) bool) int {
	sum := 0
	for _, r := range s.ranges {
		for i := r.From; i <= r.To; i++ {
			if predicate(i) {
				sum += i
			}
		}
	}
	return sum
}

func isRepeating(str string) bool {
	factors := math.FactorsOf(len(str))
	for _, f := range factors[:len(factors)-1] {
		if repeated(f, str) {
			return true
		}
	}
	return false
}

func repeated(factor int, str string) bool {
	unit := str[:factor]
	for i := 0; i+factor <= len(str); i += factor {
		if str[i:i+factor] != unit {
			return false
		}
	}
	return true
}
