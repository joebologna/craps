package point

import "strconv"

var NO_POINT = Point{0}

type Point struct {
	Value int
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

type PlayerStatus bool

const (
	NEW_PLAYER PlayerStatus = true
	CUR_PLAYER              = false
)

type PointTracker struct {
	NewPlayer PlayerStatus
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
			pt.NewPlayer = true
			pt.CurState = WIN
			pt.CurPoint = NO_POINT
		case 2:
			fallthrough
		case 3:
			fallthrough
		case 12:
			pt.NewPlayer = true
			pt.CurState = LOSE
			pt.CurPoint = NO_POINT
		default:
			pt.NewPlayer = false
			pt.CurState = POINT_SET
			pt.CurPoint = Point{roll}
		}
	} else if pt.CurState == POINT_SET {
		p := Point{roll}
		if pt.CurPoint == p {
			pt.NewPlayer = true
			pt.CurState = WIN
			pt.CurPoint = NO_POINT
		} else if roll == 7 {
			pt.NewPlayer = true
			pt.CurState = LOSE
			pt.CurPoint = NO_POINT
		}
		// else push.
	}
}

func (p PlayerStatus) String() string {
	if p {
		return "New Player"
	} else {
		return "Current Player"
	}
}
