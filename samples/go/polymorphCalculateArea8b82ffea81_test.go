package main

import (
	"testing"
	"math"
)

type Square struct {
	side float64
}

func (sq *Square) CalculateArea() float64 {
	return math.Abs(sq.side * sq.side) //area of square
}

func TestCalculateArea8b82ffea81(t *testing.T) {
	// Test case 1: Positive scenario
	sq := &Square{side: 5}
	area := sq.CalculateArea()
	expectedArea := 25.0
	if area != expectedArea {
		t.Errorf("Expected area is %.2f, but got %.2f", expectedArea, area)
	}

	// Test case 2: Zero side length
	sq = &Square{side: 0}
	area = sq.CalculateArea()
	expectedArea = 0
	if area != expectedArea {
		t.Errorf("Expected area is %.2f, but got %.2f", expectedArea, area)
	}

	// Test case 3: Negative side length
	sq = &Square{side: -5}
	area = sq.CalculateArea()
	expectedArea = 25.0 // Even though side is negative, area should be positive
	if area != expectedArea {
		t.Errorf("Expected area is %.2f, but got %.2f", expectedArea, area)
	}
}
