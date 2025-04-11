package point

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComeOutRoll(t *testing.T) {
	a := assert.New(t)

	pt := NewPointTracker()
	a.Equal(pt.CurState, COME_OUT_ROLL)
	a.Equal(pt.CurPoint, NO_POINT)

	pt.SetPoint(4)
	a.Equal(POINT_SET, pt.CurState)
	a.Equal(Point{4}, pt.CurPoint)

	pt = NewPointTracker()
	pt.SetPoint(5)
	a.Equal(POINT_SET, pt.CurState)
	a.Equal(Point{5}, pt.CurPoint)

	pt = NewPointTracker()
	pt.SetPoint(6)
	a.Equal(POINT_SET, pt.CurState)
	a.Equal(Point{6}, pt.CurPoint)

	pt = NewPointTracker()
	pt.SetPoint(8)
	a.Equal(POINT_SET, pt.CurState)
	a.Equal(Point{8}, pt.CurPoint)

	pt = NewPointTracker()
	pt.SetPoint(9)
	a.Equal(POINT_SET, pt.CurState)
	a.Equal(Point{9}, pt.CurPoint)

	pt = NewPointTracker()
	pt.SetPoint(10)
	a.Equal(POINT_SET, pt.CurState)
	a.Equal(Point{10}, pt.CurPoint)
}

func TestPointSet(t *testing.T) {
	a := assert.New(t)

	pt := NewPointTracker()
	pt.SetPoint(10)
	a.Equal(POINT_SET, pt.CurState)

	pt.SetPoint(10)
	a.Equal(WIN, pt.CurState)

	pt = NewPointTracker()
	pt.SetPoint(10)
	a.Equal(POINT_SET, pt.CurState)

	pt.SetPoint(4)
	a.Equal(POINT_SET, pt.CurState)

	pt.SetPoint(7)
	a.Equal(LOSE, pt.CurState)
}

func TestWin(t *testing.T) {
	a := assert.New(t)

	pt := NewPointTracker()
	a.Equal(pt.CurState, COME_OUT_ROLL)
	a.Equal(pt.CurPoint, NO_POINT)

	pt.SetPoint(7)
	a.Equal(pt.CurState, WIN)
	a.Equal(pt.CurPoint, NO_POINT)

	pt.SetPoint(11)
	a.Equal(pt.CurState, WIN)
	a.Equal(pt.CurPoint, NO_POINT)
}

func TestLose(t *testing.T) {
	a := assert.New(t)

	pt := NewPointTracker()
	a.Equal(pt.CurState, COME_OUT_ROLL)
	a.Equal(pt.CurPoint, NO_POINT)

	pt.SetPoint(2)
	a.Equal(pt.CurState, LOSE)
	a.Equal(pt.CurPoint, NO_POINT)

	pt.SetPoint(3)
	a.Equal(pt.CurState, LOSE)
	a.Equal(pt.CurPoint, NO_POINT)

	pt.SetPoint(12)
	a.Equal(pt.CurState, LOSE)
	a.Equal(pt.CurPoint, NO_POINT)
}
