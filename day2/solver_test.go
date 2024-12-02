package day2

import "testing"

func TestIsNotSafe(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected bool
	}{
		{
			name:     "repeating numbers",
			numbers:  []int{1, 1},
			expected: true,
		},
		{
			name:     "difference greater than 3",
			numbers:  []int{1, 5},
			expected: true,
		},
		{
			name:     "difference less than -3",
			numbers:  []int{5, 1},
			expected: true,
		},
		{
			name:     "both increasing and decreasing",
			numbers:  []int{1, 2, 1},
			expected: true,
		},
		{
			name:     "continuously increasing within range",
			numbers:  []int{1, 2, 3},
			expected: false,
		},
		{
			name:     "continuously decreasing within range",
			numbers:  []int{3, 2, 1},
			expected: false,
		},
		{
			name:     "single number",
			numbers:  []int{1},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNotSafe(tt.numbers); got != tt.expected {
				t.Errorf("isNotSafe() = %v, want %v", got, tt.expected)
			}
		})
	}
}
