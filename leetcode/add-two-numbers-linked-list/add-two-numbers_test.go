package add_two_numbers_linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			l1:   &ListNode{Val: 1, Next: nil},
			l2:   &ListNode{Val: 1, Next: nil},
			want: &ListNode{Val: 2, Next: nil},
		},
		{
			l1:   &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: nil}}},
			l2:   &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: nil}}},
			want: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: nil}}},
		},
		{
			l1:   &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 9, Next: nil}}},
			l2:   &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 5, Next: nil}}},
			want: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: nil}}}},
		},
		{
			l1:   &ListNode{Val: 9, Next: &ListNode{Val: 1, Next: &ListNode{Val: 9, Next: nil}}},
			l2:   &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 5, Next: nil}}},
			want: &ListNode{Val: 0, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: nil}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, addTwoNumbers(tt.l1, tt.l2))
		})
	}
}
