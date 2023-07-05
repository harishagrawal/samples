package main

import (
	"testing"
)

type Square struct {
	side float64
	area float64
}

func (sq *Square) CalculateArea() {
	sq.area = sq.side * sq.side //area of square
}

func TestCalculateArea(t *testing.T) {
	sq := &Square{side: 5}
	sq.CalculateArea()
	if sq.area != 25 {
		t.Errorf("CalculateArea was incorrect, got: %f, want: %f.", sq.area, 25.0)
	}
}
