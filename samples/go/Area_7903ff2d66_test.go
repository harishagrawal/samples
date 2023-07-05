package main

import (
	"testing"
)

type Shape struct {
	area float64
}

func (s *Shape) Area() float64 {
	return s.area
}

func TestArea_7903ff2d66(t *testing.T) {
	// Test case 1: Positive area
	shape1 := &Shape{area: 25.0}
	if area := shape1.Area(); area != 25.0 {
		t.Errorf("Expected area to be 25.0, but got %v", area)
	}

	// Test case 2: Zero area
	shape2 := &Shape{area: 0.0}
	if area := shape2.Area(); area != 0.0 {
		t.Errorf("Expected area to be 0.0, but got %v", area)
	}
}

func main() {
	// Running tests
	TestArea_7903ff2d66()
}
