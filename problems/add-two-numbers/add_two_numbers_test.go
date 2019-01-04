package add_two_numbers

import (
	"reflect"
	"testing"

	. "github.com/openset/leetcode/problems/000000"
)

type caseType struct {
	l1       []int
	l2       []int
	expected []int
}

func TestAddTwoNumbers(t *testing.T) {
	tests := [...]caseType{
		{
			l1:       []int{2, 4, 3},
			l2:       []int{5, 6, 4},
			expected: []int{7, 0, 8},
		},
		{
			l1:       []int{5},
			l2:       []int{5},
			expected: []int{0, 1},
		},
		{
			l1:       []int{5},
			l2:       []int{5, 9, 9},
			expected: []int{0, 0, 0, 1},
		},
	}
	for _, tc := range tests {
		l1 := SliceInt2ListNode(tc.l1)
		l2 := SliceInt2ListNode(tc.l2)
		output := ListNode2SliceInt(addTwoNumbers(l1, l2))
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.l1, tc.l2, output, tc.expected)
		}
	}
}
