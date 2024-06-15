package main

import "fmt"

func StartAddTwoNumbers() {

	//addTwoNumbers()
	// var p = Person{
	// 	FirstName: "John",
	// 	LastName:  "Snow",
	// 	Age:       45,
	// }
	//_a := twoSum([]int{15, 7, 11, 2}, 9)
	var _ListNode3 = ListNode{
		Val:  3,
		Next: nil,
	}
	var _ListNode2 = ListNode{
		Val:  4,
		Next: &_ListNode3,
	}
	var _ListNode = ListNode{
		Val:  2,
		Next: &_ListNode2,
	}

	var _ListNode6 = ListNode{
		Val:  4,
		Next: nil,
	}
	var _ListNode5 = ListNode{
		Val:  6,
		Next: &_ListNode6,
	}
	var _ListNode4 = ListNode{
		Val:  5,
		Next: &_ListNode5,
	}

	//fmt.Println(_ListNode)
	var aaa = AddTwoNumbers(&_ListNode, &_ListNode4)
	fmt.Println(aaa.Val, aaa.Next.Val, aaa.Next.Next.Val)

	//AddTwoNumbers(_ListNode.Next, _ListNode.Next)

	//var p = ListNode{"Rajeev", "Singh", 26}

}

type Person struct {
	FirstName, LastName string
	Age                 int
}
type ListNode struct {
	Next *ListNode
	Val  int
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int 0     243 564
 *     Next *ListNode nil
 * }
 */
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//head := &ListNode{0, nil}
	initialNode, sum := new(ListNode), 0
	for tempNode := initialNode; l1 != nil || l2 != nil || sum != 0; {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		tempNode.Next = &ListNode{Val: sum % 10}
		sum /= 10
		tempNode = tempNode.Next
	}
	return initialNode.Next
}
