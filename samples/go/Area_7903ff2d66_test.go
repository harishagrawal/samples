package main

import (
	"testing"
)

type Shape struct {
	area float64
}

func (sq *Shape) Area() float64 {
	return sq.area //return area of square
}

func TestArea_7903ff2d66(t *testing.T) {
	t.Run("Test with positive area value", func(t *testing.T) {
		shape := &Shape{area: 25.0}
		got := shape.Area()
		want := 25.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("Test with zero area value", func(t *testing.T) {
		shape := &Shape{area: 0.0}
		got := shape.Area()
		want := 0.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("Test with negative area value", func(t *testing.T) {
		shape := &Shape{area: -25.0}
		got := shape.Area()
		want := -25.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}
