package main

import (
	"testing"
)

type Rectangle struct {
	width, height, area float64
}

func (rect *Rectangle) CalculateArea() {
	rect.area = rect.width * rect.height //area of rectangle
}

func TestCalculateArea_acb586a213(t *testing.T) {
	// Test case 1: Positive numbers
	rect1 := &Rectangle{width: 5, height: 10}
	rect1.CalculateArea()
	if rect1.area != 50 {
		t.Error("Expected area to be 50, but got ", rect1.area)
	}

	// Test case 2: Zero value
	rect2 := &Rectangle{width: 0, height: 10}
	rect2.CalculateArea()
	if rect2.area != 0 {
		t.Error("Expected area to be 0, but got ", rect2.area)
	}

	// Test case 3: Negative numbers
	rect3 := &Rectangle{width: -5, height: -10}
	rect3.CalculateArea()
	if rect3.area != 50 {
		t.Error("Expected area to be 50, but got ", rect3.area)
	}

	// Test case 4: Mixed positive and negative numbers
	rect4 := &Rectangle{width: -5, height: 10}
	rect4.CalculateArea()
	if rect4.area != -50 {
		t.Error("Expected area to be -50, but got ", rect4.area)
	}
}
