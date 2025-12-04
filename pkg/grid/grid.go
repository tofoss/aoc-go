package grid

type Point struct {
	Y, X int
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
