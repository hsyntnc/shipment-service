package order

import (
	"reflect"
	"testing"
)

func TestFullFillOrder(t *testing.T) {
	tests := []struct {
		name                string
		quantity            int
		packSizes           []int
		expectedCombination map[int]int
		expectedTotal       int
	}{
		{
			name:                "Exact order possible",
			quantity:            32,
			expectedCombination: map[int]int{5: 1, 9: 3},
			expectedTotal:       32,
		},
		{
			name:                "Exact change is not possible, find closest.",
			quantity:            13,
			expectedCombination: map[int]int{5: 1, 9: 1},
			expectedTotal:       14,
		},
		{
			name:                "With thousands of orders",
			quantity:            500000,
			expectedCombination: map[int]int{9: 55555, 5: 1},
			expectedTotal:       500000,
		},
		{
			name:                "With 1 single order",
			quantity:            1,
			expectedCombination: map[int]int{5: 1},
			expectedTotal:       5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			totalAmount, packs := FullFillOrder(test.quantity)
			if totalAmount != test.expectedTotal || !reflect.DeepEqual(packs, test.expectedCombination) {
				t.Errorf("Test %s failed: expected order amount is %d and %v packs, got %d total and %v packs",
					test.name, test.expectedTotal, test.expectedCombination, totalAmount, packs)
			}
		})
	}
}
