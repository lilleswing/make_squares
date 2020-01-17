package make_squares

import "strconv"

type NumberRequest struct {
	Value int
}

func SquareIt(a int) string {
	return strconv.Itoa(a * a)
}
