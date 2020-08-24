/* 155.Min Stack

References:
    https://leetcode.com/problems/min-stack/
    https://leetcode-cn.com/problems/min-stack/

*/

package leetcode

// MinStack 辅助栈
//
// 定义「数据栈」支持`push`、`pop`、`top` 操作
// 定义「辅助栈」保持栈顶元素为当前的最小值
//
// 当元素入数据栈时，如果此元素不大于辅助栈的栈顶元素，那么此元素也需要压入辅助栈的栈顶；
// 当元素从栈顶被弹出时，如果此元素等于辅助栈的栈顶元素，即此元素就是当前的最小值，也需要从辅助栈弹出
//
// * 时间复杂度: O(1)
// * 空间复杂度: O(N)，其中`N`为总操作数，最坏情况下，我们会连续插入`N`个元素，此时两个栈占用的空间为 O(N)
//
type MinStack struct {
	stack    []int // the stack to keep elements
	minStack []int // the stack to keep min element
}

// Constructor returns new MinStack.
func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (s *MinStack) Push(val int) {
	s.stack = append(s.stack, val)
	if len(s.minStack) == 0 || val <= s.minStack[len(s.minStack)-1] {
		s.minStack = append(s.minStack, val)
	}
}

func (s *MinStack) Pop() {
	val := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	if s.minStack[len(s.minStack)-1] == val {
		s.minStack = s.minStack[:len(s.minStack)-1]
	}
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	return s.minStack[len(s.minStack)-1]
}
