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

type BetTracker struct {
	NewPlayer PlayerStatus
	CurState  PointState
	CurPoint  Point
	PassBet   bool
}

func NewBetTracker() *BetTracker {
	return &BetTracker{true, COME_OUT_ROLL, NO_POINT, true}
}
func (bt *BetTracker) SetPassBet(yesNo bool) {
	bt.PassBet = yesNo
}

func (bt *BetTracker) WinLose(won bool) {
	if bt.PassBet {
		if won {
			bt.CurState = WIN
		} else {
			bt.CurState = LOSE
		}
	} else {
		if !won {
			bt.CurState = LOSE
		} else {
			bt.CurState = WIN
		}
	}
}

func (bt *BetTracker) SetPoint(roll int) {
	if bt.CurState == COME_OUT_ROLL {
		switch roll {
		case 7:
			fallthrough
		case 11:
			bt.NewPlayer = true
			bt.WinLose(true)
			bt.CurPoint = NO_POINT
		case 2:
			fallthrough
		case 3:
			fallthrough
		case 12:
			bt.NewPlayer = true
			bt.WinLose(false)
			bt.CurPoint = NO_POINT
		default:
			bt.NewPlayer = false
			bt.CurState = POINT_SET
			bt.CurPoint = Point{roll}
		}
	} else if bt.CurState == POINT_SET {
		p := Point{roll}
		if bt.CurPoint == p {
			bt.NewPlayer = true
			bt.WinLose(true)
			bt.CurPoint = NO_POINT
		} else if roll == 7 {
			bt.NewPlayer = true
			bt.WinLose(false)
			bt.CurPoint = NO_POINT
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
