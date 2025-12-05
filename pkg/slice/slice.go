package slice

import (
	"fmt"
	"strconv"

	"golang.org/x/exp/constraints"
)

func MapErr[T, R any](slice []T, transform func(T) (R, error)) ([]R, error) {
	res := make([]R, len(slice))
	for i := range slice {
		r, err := transform(slice[i])
		if err != nil {
			return nil, err
		}
		res[i] = r
	}
	return res, nil
}

func Map[T, R any](slice []T, transform func(T) R) []R {
	res := make([]R, len(slice))
	for i := range slice {
		r := transform(slice[i])
		res[i] = r
	}
	return res
}

func SumOfStrs(strInts []string) (int, error) {
	ints, err := MapErr(strInts, strconv.Atoi)
	if err != nil {
		return 0, fmt.Errorf("cannot convert strings to int")
	}
	return Sum(ints), nil
}

func Sum[T constraints.Integer | constraints.Float](numbers []T) T {
	var sum T
	for _, n := range numbers {
		sum += n
	}
	return sum
}
