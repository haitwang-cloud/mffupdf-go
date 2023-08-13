package Int

// https://leetcode.cn/problems/remove-element/?envType=study-plan-v2&envId=top-interview-150
func removeElement(nums []int, val int) int {

	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
