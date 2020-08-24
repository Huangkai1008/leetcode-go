/* 21.Merge Two Sorted Lists

References:
    https://leetcode.com/problems/merge-two-sorted-lists/
    https://leetcode-cn.com/problems/merge-two-sorted-lists/

*/

package leetcode

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeTwoLists1 迭代
//
// 判断哪一个链表的头节点的值更小，将较小值的节点添加到结果里
// 当一个节点被添加到结果里之后，将对应链表中的节点向后移一位
//
// 增加一个「虚拟头节点(`dummy node`)」 / 「哨兵节点(sentinel node)」
// 可以使返回合并后的结果更加简单
//
// * 时间复杂度: O(n+m)，其中 `n` 和 `m` 分别为两个链表的长度
// * 空间复杂度: O(1)
//
func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}

	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}

	return dummy.Next
}
