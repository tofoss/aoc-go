package slice

import (
	"fmt"
	"strconv"
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

func Sum(ints []int) int {
	sum := 0
	for _, n := range ints {
		sum += n
	}
	return sum
}
