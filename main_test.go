package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	// // Arrange
	// numbers := []int{1, 2, 5, 7, -2, -7, 15, 3}
	// expected := 15
	// // Act
	// result := max(numbers)
	// // Assert
	// if result != expected {
	// 	t.Errorf("max() = %v, want %v", result, expected)
	// }

	// Arrange
	// Табличное тестирование (table-driven tests).
	// Создается переменная tests , ее тип - «срез из анонимных структур,
	// содержащих поля name, numbers и want
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		// Именованные входные параметры для целевой функции.
		numbers []int
		want    int
	}{
		{"good", []int{1, 2, 5, 7, -2, -7, 15, 3}, 15},
		{"empty", []int{}, 0},
	}

	// Act
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxnumber(tt.numbers)
			t.Logf("Calling maxnumber(%v), result %d\n)", tt.numbers, got)
			// Assert
			assert.Equal(t, tt.want, got, fmt.Sprintf("max() = %v, want %v", got, tt.want))
			// if got != tt.want {
			// 	t.Errorf("max() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func TestMaxIndex(t *testing.T) {
	// Arrange
	// Табличное тестирование (table-driven tests).
	// Создается переменная tests , ее тип - «срез из анонимных структур,
	// содержащих поля name, numbers и want
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		numbers []int
		want    int
	}{
		{"good", []int{6, 2, 5, 7, -2, -7, 15, 3}, 6},
		{"empty", []int{}, 0},
	}

	// Act
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxindex(tt.numbers)
			t.Logf("Calling maxindex(%v), result %d\n)", tt.numbers, got)
			// Assert
			assert.Equal(t, tt.want, got, fmt.Sprintf("maxindex() = %v, want %v", got, tt.want))
			// if got != tt.want {
			// 	t.Errorf("maxIndex() = %v, want %v", got, tt.want)
			// }
		})
	}
}
