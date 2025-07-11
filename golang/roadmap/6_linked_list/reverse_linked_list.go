package linked_list

import datastruc "github.com/ViniNepo/algorithm/golang/date_structure/arrays_and_hashing"

/*
Reverse Linked List
Given the beginning of a singly linked list head, reverse the list, and return the new beginning of the list.

Example 1:

Input: head = [0,1,2,3]

Output: [3,2,1,0]
Example 2:

Input: head = []

Output: []
*/

func ReverseLinkedList(head *datastruc.ListNode) *datastruc.ListNode {
	var prev *datastruc.ListNode
	curr := head

	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	return prev
}
