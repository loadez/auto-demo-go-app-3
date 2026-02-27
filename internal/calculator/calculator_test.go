package calculator_test

import (
	"testing"

	"github.com/marcosfilipe/auto-demo-go-app/internal/calculator"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, 1, 0},
		{100, 200, 300},
	}

	for _, tt := range tests {
		got := calculator.Add(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{5, 3, 2},
		{0, 0, 0},
		{-1, -1, 0},
		{10, 20, -10},
	}

	for _, tt := range tests {
		got := calculator.Subtract(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Subtract(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{2, 3, 6},
		{0, 5, 0},
		{-2, 3, -6},
		{-2, -3, 6},
	}

	for _, tt := range tests {
		got := calculator.Multiply(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Multiply(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestDivide(t *testing.T) {
	got, err := calculator.Divide(10, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 5 {
		t.Errorf("Divide(10, 2) = %d, want 5", got)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := calculator.Divide(10, 0)
	if err != calculator.ErrDivideByZero {
		t.Errorf("Divide(10, 0) error = %v, want ErrDivideByZero", err)
	}
}
