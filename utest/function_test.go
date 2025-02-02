package utest

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b   int
		result int
	}{
		{1, 2, 3},
		{-1, -2, -3},
		{0, 0, 0},
		{100, 200, 300},
	}

	for _, tt := range tests {
		t.Run("Add", func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.result {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.result)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		a, b     int
		result   int
		hasError bool
	}{
		{10, 2, 5, false},
		{7, 0, 0, true},
		{9, 3, 3, false},
		{-10, -2, 5, false},
	}

	for _, tt := range tests {
		t.Run("Divide", func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)
			if tt.hasError {
				if err == nil {
					t.Errorf("Divide(%d, %d) = %d; want error", tt.a, tt.b, got)
				}
			} else {
				if err != nil {
					t.Errorf("Divide(%d, %d) = %d, %v; want %d, nil", tt.a, tt.b, got, err, tt.result)
				} else if got != tt.result {
					t.Errorf("Divide(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.result)
				}
			}
		})
	}
}
