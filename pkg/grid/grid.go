package grid

type Point struct {
	Y, X int
}

func (p *Point) Down() Point {
	return Point{p.Y + Adjecent.Bottom.Y, p.X + Adjecent.Bottom.X}
}

func (p *Point) DownLeft() Point {
	return Point{p.Y + Adjecent.BottomLeft.Y, p.X + Adjecent.BottomLeft.X}
}

func (p *Point) DownRight() Point {
	return Point{p.Y + Adjecent.BottomRight.Y, p.X + Adjecent.BottomRight.X}
}

func OutOfBounds[T any](point Point, matrix [][]T) bool {
	rows := len(matrix) - 1
	columns := len(matrix[0]) - 1
	return point.OffGrid(rows, columns)
}

func (p *Point) OffGrid(rows, columns int) bool {
	if rows < 1 || columns < 1 {
		return true
	}
	if p.Y < 0 || p.X < 0 || p.Y > rows || p.X > columns {
		return true
	}
	return false
}

var Adjecent = struct {
	TopLeft     Point
	Top         Point
	TopRight    Point
	Right       Point
	BottomRight Point
	Bottom      Point
	BottomLeft  Point
	Left        Point
}{
	TopLeft:     Point{Y: -1, X: -1},
	Top:         Point{Y: -1, X: 0},
	TopRight:    Point{Y: -1, X: 1},
	Right:       Point{Y: 0, X: 1},
	BottomRight: Point{Y: 1, X: 1},
	Bottom:      Point{Y: 1, X: 0},
	BottomLeft:  Point{Y: 1, X: -1},
	Left:        Point{Y: 0, X: -1},
}

var Adjencencies = []Point{
	Adjecent.TopLeft,
	Adjecent.Top,
	Adjecent.TopRight,
	Adjecent.Right,
	Adjecent.BottomRight,
	Adjecent.Bottom,
	Adjecent.BottomLeft,
	Adjecent.Left,
}
