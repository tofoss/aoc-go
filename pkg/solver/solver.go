package solver

type Solver interface {
	Part1() (string, error)
	Part2() (string, error)
}
