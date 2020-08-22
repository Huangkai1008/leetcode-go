/* 20.Valid Parentheses

References:
    https://leetcode.com/problems/valid-parentheses/
    https://leetcode-cn.com/problems/valid-parentheses/

*/

package leetcode

// isValid1 栈
//
// 1.先判断字符串的长度是否是偶数，如果不是，肯定无法匹配；
// 2.维护一个栈，遍历给定的字符串 `s` ：
//   a.当遇到左括号时入栈
//   b.当遇到右括号时需要去栈顶找匹配的左括号，如果此时栈中没有元素或者栈顶对应的左括号类型不匹配则返回False
//     如果存在匹配则将栈顶元素出栈
// 3.遍历结束后如果栈中还存在元素，那么说明没有一一匹配上，返回False
//
// * 时间复杂度: O(N)，其中N为字符串的长度
// * 空间复杂度: O(N)，其中N为字符串的长度，最坏情况下栈的长度可能等于字符串的长度
//
func isValid1(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
		char := s[i]
		if pair, exist := pairs[char]; exist {
			if len(stack) == 0 || stack[len(stack)-1] != pair {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}
