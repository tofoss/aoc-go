package registry

import "github.com/tofoss/aoc-go/pkg/solver"

var Registry = make(map[int]map[int]func(string) (solver.Solver, error))

func Register(year, day int, factory func(string) (solver.Solver, error)) {
	if Registry[year] == nil {
		Registry[year] = make(map[int]func(string) (solver.Solver, error))
	}
	Registry[year][day] = factory
}
