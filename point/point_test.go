package point

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComeOutRoll(t *testing.T) {
	a := assert.New(t)

	pt := NewPointTracker()
	a.Equal("New Player", pt.NewPlayer.String())
	a.Equal(COME_OUT_ROLL, pt.CurState)
	a.Equal(NO_POINT, pt.CurPoint)

	pt.SetPoint(4)
	a.Equal("Current Player", pt.NewPlayer.String())
	a.Equal(POINT_SET, pt.CurState)
	a.Equal(Point{4}, pt.CurPoint)

	pt = NewPointTracker()
	pt.SetPoint(5)
	a.Equal("Current Player", pt.NewPlayer.String())
	a.Equal(POINT_SET, pt.CurState)
	a.Equal(Point{5}, pt.CurPoint)

	pt = NewPointTracker()
	pt.SetPoint(6)
	a.Equal("Current Player", pt.NewPlayer.String())
	a.Equal(POINT_SET, pt.CurState)
	a.Equal(Point{6}, pt.CurPoint)

	pt = NewPointTracker()
	pt.SetPoint(8)
	a.Equal("Current Player", pt.NewPlayer.String())
	a.Equal(POINT_SET, pt.CurState)
	a.Equal(Point{8}, pt.CurPoint)

	pt = NewPointTracker()
	pt.SetPoint(9)
	a.Equal("Current Player", pt.NewPlayer.String())
	a.Equal(POINT_SET, pt.CurState)
	a.Equal(Point{9}, pt.CurPoint)

	pt = NewPointTracker()
	pt.SetPoint(10)
	a.Equal("Current Player", pt.NewPlayer.String())
	a.Equal(POINT_SET, pt.CurState)
	a.Equal(Point{10}, pt.CurPoint)
}

func TestWin(t *testing.T) {
	a := assert.New(t)

	pt := NewPointTracker()
	a.Equal("New Player", pt.NewPlayer.String())
	a.Equal(pt.CurState, COME_OUT_ROLL)
	a.Equal(pt.CurPoint, NO_POINT)

	pt.SetPoint(7)
	a.Equal("New Player", pt.NewPlayer.String())
	a.Equal(pt.CurState, WIN)
	a.Equal(pt.CurPoint, NO_POINT)

	pt.SetPoint(11)
	a.Equal("New Player", pt.NewPlayer.String())
	a.Equal(pt.CurState, WIN)
	a.Equal(pt.CurPoint, NO_POINT)
}

func TestLose(t *testing.T) {
	a := assert.New(t)

	pt := NewPointTracker()
	a.Equal("New Player", pt.NewPlayer.String())
	a.Equal(pt.CurState, COME_OUT_ROLL)
	a.Equal(pt.CurPoint, NO_POINT)

	pt.SetPoint(2)
	a.Equal("New Player", pt.NewPlayer.String())
	a.Equal(pt.CurState, LOSE)
	a.Equal(pt.CurPoint, NO_POINT)

	pt.SetPoint(3)
	a.Equal("New Player", pt.NewPlayer.String())
	a.Equal(pt.CurState, LOSE)
	a.Equal(pt.CurPoint, NO_POINT)

	pt.SetPoint(12)
	a.Equal("New Player", pt.NewPlayer.String())
	a.Equal(pt.CurState, LOSE)
	a.Equal(pt.CurPoint, NO_POINT)
}

func TestString(t *testing.T) {
	a := assert.New(t)

	pt := NewPointTracker()
	a.Equal("Come out Roll", pt.CurState.String())

	pt.SetPoint(4)
	a.Equal("Point Set", pt.CurState.String())

	pt.SetPoint(4)
	a.Equal("Win", pt.CurState.String())

	pt = NewPointTracker()
	pt.SetPoint(4)
	pt.SetPoint(7)
	a.Equal("Lose", pt.CurState.String())

	pt.CurState = LOSE + 1
	a.Equal("Unknown", pt.CurState.String())
}

func TestPointString(t *testing.T) {
	a := assert.New(t)

	pt := NewPointTracker()
	a.Equal("No Point", pt.CurPoint.String())
	pt.SetPoint(4)
	a.Equal("4", pt.CurPoint.String())
}
