package day05

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/tofoss/aoc-go/pkg/registry"
	"github.com/tofoss/aoc-go/pkg/slice"
	"github.com/tofoss/aoc-go/pkg/solver"
)

const year = 2025
const day = 5

func init() {
	registry.Register(year, day, func(input []string) (solver.Solver, error) {
		return New(input)
	})
}

type Solution struct {
	db database
}

type Range struct {
	min, max int64
}

type database struct {
	FreshRanges []Range
	Ingredients []int64
}

func New(input []string) (*Solution, error) {
	readRanges := true
	db := database{}
	for _, line := range input {
		if line == "" {
			readRanges = false
			continue
		}
		if readRanges {
			parts := strings.Split(line, "-")
			min, err := strconv.ParseInt(parts[0], 10, 64)
			max, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				return nil, err
			}
			db.FreshRanges = append(db.FreshRanges, Range{min, max})
		} else {
			ing, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				return nil, err
			}
			db.Ingredients = append(db.Ingredients, ing)
		}
	}

	return &Solution{db}, nil
}

func (s *Solution) Part1() (string, error) {
	freshIngredients := 0
	for _, ingredient := range s.db.Ingredients {
		for _, freshRange := range s.db.FreshRanges {
			if ingredient >= freshRange.min && ingredient <= freshRange.max {
				freshIngredients++
				break
			}
		}
	}
	return fmt.Sprintf("%d", freshIngredients), nil
}

func mergeNextOverlap(ranges []Range) []Range {
	for i := range ranges {
		for j := range ranges {
			if i == j {
				continue
			}
			if ranges[i].max >= ranges[j].min && ranges[j].max >= ranges[i].min {
				newRange := Range{
					min: min(ranges[i].min, ranges[j].min),
					max: max(ranges[i].max, ranges[j].max),
				}
				if i > j {
					i, j = j, i
				}
				ranges = slices.Delete(ranges, j, j+1)
				ranges = slices.Delete(ranges, i, i+1)
				return append(ranges, newRange)
			}
		}
	}
	return ranges
}

func (s *Solution) Part2() (string, error) {
	merged := s.db.FreshRanges

	for {
		next := mergeNextOverlap(merged)
		if len(merged) == len(next) {
			break
		}
		merged = next
	}

	counts := slice.Map(merged, func(r Range) int64 {
		return 1 + r.max - r.min
	})

	return fmt.Sprintf("%d", slice.Sum(counts)), nil
}
