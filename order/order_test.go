package order

import (
	"reflect"
	"testing"
)

// packSizes = [100, 200, 500]
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
			quantity:            300,
			expectedCombination: map[int]int{100: 1, 200: 1},
			expectedTotal:       300,
		},
		{
			name:                "Exact change is not possible, find closest.",
			quantity:            150,
			expectedCombination: map[int]int{200: 1},
			expectedTotal:       200,
		},
		{
			name:                "With thousands of orders",
			quantity:            500000,
			expectedCombination: map[int]int{500: 1000},
			expectedTotal:       500000,
		},
		{
			name:                "With 1 single order",
			quantity:            1,
			expectedCombination: map[int]int{100: 1},
			expectedTotal:       100,
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
