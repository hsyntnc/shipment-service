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
			packSizes:           []int{5, 7, 9},
			expectedCombination: map[int]int{5: 1, 9: 3},
			expectedTotal:       32,
		},
		{
			name:                "Exact change is not possible, find closest.",
			quantity:            13,
			packSizes:           []int{5, 7, 9},
			expectedCombination: map[int]int{5: 1, 9: 1},
			expectedTotal:       14,
		},
		{
			name:                "Exact change, single pack option",
			quantity:            26,
			packSizes:           []int{7},
			expectedCombination: map[int]int{7: 4},
			expectedTotal:       28,
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
