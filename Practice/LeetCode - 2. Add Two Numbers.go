package main

func StartAddTwoNumbers() {
	//myLikedList := ll.New()
	//print(myLikedList) // nil
	// _a := addTwoNumbers({2, 4, 3}, {2, 4, 3})
	// fmt.Println(_a)
	//addTwoNumbers()
	//_a := ListNode();
	// var p ListNode
	// p = ListNode{5, 4}

	//var p = ListNode{"Rajeev", "Singh", 26}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	current := head
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		current.Next = new(ListNode)
		current.Next.Val = sum % 10
		current = current.Next
	}
	return head.Next
}
