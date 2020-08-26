/*232. Implement Queue using Stacks

References:
    https://leetcode.com/problems/implement-queue-using-stacks/
    https://leetcode-cn.com/problems/implement-queue-using-stacks/

*/

package leetcode

// MyQueue 双栈
//
// 定义「输入栈」保存入队（`push`）的元素；
// 每次 `pop` 或 `peek` 时，当「输出栈」为空时，输入栈的元素全部倒序压入「输出栈」中
//
// * 时间复杂度: `push` 和 `empty` 为 O(1), `pop` 和 `peek` 为均摊 O(1)
// * 空间复杂度: O(N)
//
type MyQueue struct {
	inStack  []int
	outStack []int
}

// Constructor returns new MyQueue.
func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.inStack = append(q.inStack, x)
}

func (q *MyQueue) Pop() int {
	if len(q.outStack) == 0 {
		for len(q.inStack) > 0 {
			q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
			q.inStack = q.inStack[:len(q.inStack)-1]
		}
	}
	s := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]
	return s
}

func (q *MyQueue) Peek() int {
	if len(q.outStack) == 0 {
		for len(q.inStack) > 0 {
			q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
			q.inStack = q.inStack[:len(q.inStack)-1]
		}
	}
	return q.outStack[len(q.outStack)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.inStack) == 0 && len(q.outStack) == 0
}
