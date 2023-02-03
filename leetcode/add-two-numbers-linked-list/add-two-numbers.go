package add_two_numbers_linked_list

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.



Example 1:


Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]
Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]


Constraints:

The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func value(l *ListNode) int {
	if l == nil {
		return 0
	}

	return l.Val
}

type LinkedList struct {
	head *ListNode
	tail *ListNode
}

func (l *LinkedList) AdLast(value int) {
	newNode := &ListNode{
		Val:  value,
		Next: nil,
	}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	l.tail.Next = newNode
	l.tail = l.tail.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	listOne := LinkedList{l1, nil}
	listTwo := LinkedList{l2, nil}
	result := LinkedList{}

	currentListOne := listOne.head
	currentListTwo := listTwo.head

	carry := 0
	for {
		value := value(currentListOne) + value(currentListTwo) + carry
		carry = 0

		if value > 9 {
			value = value % 10
			carry = 1
		}

		result.AdLast(value)

		if currentListOne != nil {
			currentListOne = currentListOne.Next
		}

		if currentListTwo != nil {
			currentListTwo = currentListTwo.Next
		}

		if currentListOne == nil && currentListTwo == nil && carry == 0 {
			break
		}
	}

	return result.head
}
