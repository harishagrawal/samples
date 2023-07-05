package main

import (
	"math"
	"testing"
)

type Square struct {
	side float64
	area float64
}

func (sq *Square) CalculateArea() {
	sq.area = math.Abs(sq.side) * math.Abs(sq.side) //area of square
}

func TestCalculateArea_8b82ffea81(t *testing.T) {
	// Test case 1: Normal scenario
	sq1 := &Square{side: 5}
	sq1.CalculateArea()
	if sq1.area != 25 {
		t.Errorf("Area of square is incorrect, got: %f, want: %f.", sq1.area, 25.0)
	}

	// Test case 2: Edge case with side 0
	sq2 := &Square{side: 0}
	sq2.CalculateArea()
	if sq2.area != 0 {
		t.Errorf("Area of square is incorrect, got: %f, want: %f.", sq2.area, 0.0)
	}

	// Test case 3: Test with negative side
	sq3 := &Square{side: -5}
	sq3.CalculateArea()
	if sq3.area != 25 {
		t.Errorf("Area of square is incorrect, got: %f, want: %f.", sq3.area, 25.0)
	}
}
