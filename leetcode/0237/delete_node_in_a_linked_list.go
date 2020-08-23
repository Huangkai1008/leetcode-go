/*237.Delete Node in a Linked List

References:
    https://leetcode.com/problems/delete-node-in-a-linked-list/
    https://leetcode-cn.com/problems/delete-node-in-a-linked-list/

*/

package leetcode

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// deleteNode1
//
// 将后一节点的值复制给待删除节点，
// 然后将后一节点当作「待删除节点」来进行删除
//
// * 时间复杂度: O(1)
// * 空间复杂度: O(1)
//
func deleteNode1(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
