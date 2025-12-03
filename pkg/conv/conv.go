package conv

import "fmt"

func RuneToInt(r rune) (int, error) {
	if r >= '0' && r <= '9' {
		return int(r - '0'), nil
	} else {
		return 0, fmt.Errorf("cannot convert rune %v to int", r)
	}
}
