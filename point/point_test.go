package point

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComeOutRoll(t *testing.T) {
	a := assert.New(t)

	bt := NewBetTracker()
	a.Equal("New Player", bt.NewPlayer.String())
	a.Equal(COME_OUT_ROLL, bt.CurState)
	a.Equal(NO_POINT, bt.CurPoint)

	bt.SetPoint(4)
	a.Equal("Current Player", bt.NewPlayer.String())
	a.Equal(POINT_SET, bt.CurState)
	a.Equal(Point{4}, bt.CurPoint)

	bt = NewBetTracker()
	bt.SetPoint(5)
	a.Equal("Current Player", bt.NewPlayer.String())
	a.Equal(POINT_SET, bt.CurState)
	a.Equal(Point{5}, bt.CurPoint)

	bt = NewBetTracker()
	bt.SetPoint(6)
	a.Equal("Current Player", bt.NewPlayer.String())
	a.Equal(POINT_SET, bt.CurState)
	a.Equal(Point{6}, bt.CurPoint)

	bt = NewBetTracker()
	bt.SetPoint(8)
	a.Equal("Current Player", bt.NewPlayer.String())
	a.Equal(POINT_SET, bt.CurState)
	a.Equal(Point{8}, bt.CurPoint)

	bt = NewBetTracker()
	bt.SetPoint(9)
	a.Equal("Current Player", bt.NewPlayer.String())
	a.Equal(POINT_SET, bt.CurState)
	a.Equal(Point{9}, bt.CurPoint)

	bt = NewBetTracker()
	bt.SetPoint(10)
	a.Equal("Current Player", bt.NewPlayer.String())
	a.Equal(POINT_SET, bt.CurState)
	a.Equal(Point{10}, bt.CurPoint)
}

func TestWinPassBet(t *testing.T) {
	a := assert.New(t)

	bt := NewBetTracker()
	a.Equal("New Player", bt.NewPlayer.String())
	a.Equal(bt.CurState, COME_OUT_ROLL)
	a.Equal(bt.CurPoint, NO_POINT)

	bt.SetPoint(7)
	a.Equal("New Player", bt.NewPlayer.String())
	a.Equal(bt.CurState, WIN)
	a.Equal(bt.CurPoint, NO_POINT)

	bt.SetPoint(11)
	a.Equal("New Player", bt.NewPlayer.String())
	a.Equal(bt.CurState, WIN)
	a.Equal(bt.CurPoint, NO_POINT)
}

func TestLosePassBet(t *testing.T) {
	a := assert.New(t)

	bt := NewBetTracker()
	a.Equal("New Player", bt.NewPlayer.String())
	a.Equal(bt.CurState, COME_OUT_ROLL)
	a.Equal(bt.CurPoint, NO_POINT)

	bt.SetPoint(2)
	a.Equal("New Player", bt.NewPlayer.String())
	a.Equal(bt.CurState, LOSE)
	a.Equal(bt.CurPoint, NO_POINT)

	bt.SetPoint(3)
	a.Equal("New Player", bt.NewPlayer.String())
	a.Equal(bt.CurState, LOSE)
	a.Equal(bt.CurPoint, NO_POINT)

	bt.SetPoint(12)
	a.Equal("New Player", bt.NewPlayer.String())
	a.Equal(bt.CurState, LOSE)
	a.Equal(bt.CurPoint, NO_POINT)
}

func TestWinNoPassBet(t *testing.T) {
	a := assert.New(t)

	bt := NewBetTracker()
	bt.SetPassBet(false)
	a.Equal("New Player", bt.NewPlayer.String())
	a.Equal(bt.CurState, COME_OUT_ROLL)
	a.Equal(bt.CurPoint, NO_POINT)

	bt.SetPoint(7)
	a.Equal("New Player", bt.NewPlayer.String())
	a.Equal(bt.CurState, WIN)
	a.Equal(bt.CurPoint, NO_POINT)

	bt.SetPoint(11)
	a.Equal("New Player", bt.NewPlayer.String())
	a.Equal(bt.CurState, WIN)
	a.Equal(bt.CurPoint, NO_POINT)
}

func TestLoseNoPassBet(t *testing.T) {
	a := assert.New(t)

	bt := NewBetTracker()
	bt.SetPassBet(false)
	a.Equal("New Player", bt.NewPlayer.String())
	a.Equal(bt.CurState, COME_OUT_ROLL)
	a.Equal(bt.CurPoint, NO_POINT)

	bt.SetPoint(2)
	a.Equal("New Player", bt.NewPlayer.String())
	a.Equal(bt.CurState, LOSE)
	a.Equal(bt.CurPoint, NO_POINT)

	bt.SetPoint(3)
	a.Equal("New Player", bt.NewPlayer.String())
	a.Equal(bt.CurState, LOSE)
	a.Equal(bt.CurPoint, NO_POINT)

	bt.SetPoint(12)
	a.Equal("New Player", bt.NewPlayer.String())
	a.Equal(bt.CurState, LOSE)
	a.Equal(bt.CurPoint, NO_POINT)
}

func TestString(t *testing.T) {
	a := assert.New(t)

	bt := NewBetTracker()
	a.Equal("Come out Roll", bt.CurState.String())

	bt.SetPoint(4)
	a.Equal("Point Set", bt.CurState.String())

	bt.SetPoint(4)
	a.Equal("Win", bt.CurState.String())

	bt = NewBetTracker()
	bt.SetPoint(4)
	bt.SetPoint(7)
	a.Equal("Lose", bt.CurState.String())

	bt.CurState = LOSE + 1
	a.Equal("Unknown", bt.CurState.String())
}

func TestPointString(t *testing.T) {
	a := assert.New(t)

	bt := NewBetTracker()
	a.Equal("No Point", bt.CurPoint.String())
	bt.SetPoint(4)
	a.Equal("4", bt.CurPoint.String())
}
