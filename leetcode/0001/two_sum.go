/* 1.Two Sum

References: https://leetcode.com/problems/two-sum/

*/

package leetcode

// twoSum1 暴力枚举
//
// 枚举数组中的每一个数 `x` 和其后的每一个数 `y` ，当`x`和`y`之和满足条件时即返回
//
// * 时间复杂度: O(N^2)，其中N为数组中的元素数量
// * 空间复杂度: O(1)
//
func twoSum1(nums []int, target int) []int {
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			if v+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// twoSum2 哈希表
// 创建一个哈希表，遍历数组中的每一个数 `x`
// 判断 `target - x` 是否在哈希表中：
// 	   1. 如果在那么直接返回；
//     2. 如果不存在将 `x` 的值作为key, `x`在数组中的位置 `index` 作为value存入哈希表中
// 这样可以将每个找 `target - x` 的时间复杂度降到O(1)级别
//
// * 时间复杂度: O(N)，其中N为数组中的元素数量
// * 空间复杂度: O(N)，其中N为数组中的元素数量，主要为创建哈希表的开销
//
func twoSum2(nums []int, target int) []int {
	hashTable := map[int]int{}
	for index, num := range nums {
		if pos, exist := hashTable[target-num]; exist {
			return []int{pos, index}
		}
		hashTable[num] = index
	}
	return nil
}
