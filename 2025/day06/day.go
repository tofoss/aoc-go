package day06

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
const day = 6

func init() {
	registry.Register(year, day, func(input []string) (solver.Solver, error) {
		return New(input)
	})
}

type worksheet struct {
	humanColumns [][]int
	cephColumns  [][]int
	operations   []string
}

type Solution struct {
	sheet worksheet
}

func New(input []string) (*Solution, error) {
	operations := strings.Fields(input[len(input)-1])
	input = input[:len(input)-1]
	humanColumns := make([][]int, len(operations))

	for i := range input {
		for j, str := range strings.Fields(input[i]) {
			n, err := strconv.Atoi(str)
			if err != nil {
				return nil, err
			}
			humanColumns[j] = append(humanColumns[j], n)
		}
	}

	byteMatrixRotated := make([][]byte, len(input[0]))
	for i := range byteMatrixRotated {
		byteMatrixRotated[i] = make([]byte, len(input))
		for j := range byteMatrixRotated[i] {
			byteMatrixRotated[i][j] = input[len(input)-1-j][i]
		}
		slices.Reverse(byteMatrixRotated[i])
	}

	cephColumns := [][]int{}
	curLine := []int{}

	for i := range byteMatrixRotated {
		bytestr := strings.TrimSpace(string(byteMatrixRotated[i]))
		if bytestr == "" {
			cephColumns = append(cephColumns, curLine)
			curLine = []int{}
		} else {
			n, err := strconv.Atoi(bytestr)
			if err != nil {
				return nil, err
			}
			curLine = append(curLine, n)
		}
	}
	cephColumns = append(cephColumns, curLine)

	return &Solution{worksheet{humanColumns, cephColumns, operations}}, nil
}

func (s *Solution) Part1() (string, error) {
	total := 0
	for i := range s.sheet.humanColumns {
		if s.sheet.operations[i] == "+" {
			total += slice.Sum(s.sheet.humanColumns[i])
		} else {
			total += slice.Multiply(s.sheet.humanColumns[i])
		}
	}
	return fmt.Sprintf("%d", total), nil
}

func (s *Solution) Part2() (string, error) {
	total := 0
	for i := range s.sheet.cephColumns {
		if s.sheet.operations[i] == "+" {
			total += slice.Sum(s.sheet.cephColumns[i])
		} else {
			total += slice.Multiply(s.sheet.cephColumns[i])
		}
	}
	return fmt.Sprintf("%d", total), nil
}
