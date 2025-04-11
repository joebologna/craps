package point

import "strconv"

var NO_POINT = Point{0}

type Point struct {
	Value int
}

func (p *Point) SetPoint(newPoint int) {
	p.Value = newPoint
}

func (p Point) String() string {
	if p == NO_POINT {
		return "No Point"
	}
	return strconv.FormatInt(int64(p.Value), 10)
}

type PointState int

const (
	COME_OUT_ROLL PointState = iota
	POINT_SET
	WIN
	LOSE
)

func (p PointState) String() string {
	switch p {
	case COME_OUT_ROLL:
		return "Come out Roll"
	case POINT_SET:
		return "Point Set"
	case WIN:
		return "Win"
	case LOSE:
		return "Lose"
	default:
		return "Unknown"
	}
}

type PointTracker struct {
	NewPlayer bool
	CurState  PointState
	CurPoint  Point
}

func NewPointTracker() *PointTracker {
	return &PointTracker{true, COME_OUT_ROLL, NO_POINT}
}

func (pt *PointTracker) SetPoint(roll int) {
	if pt.CurState == COME_OUT_ROLL {
		switch roll {
		case 7:
			fallthrough
		case 11:
			pt.CurState = WIN
		case 2:
			fallthrough
		case 3:
			fallthrough
		case 12:
			pt.CurState = LOSE
			pt.CurPoint = NO_POINT
		default:
			pt.CurState = POINT_SET
			pt.CurPoint = Point{roll}
		}
	} else if pt.CurState == POINT_SET {
		p := Point{roll}
		if pt.CurPoint == p {
			pt.CurState = WIN
		} else if roll == 7 {
			pt.CurState = LOSE
		}
		// else push.
	}
}
