package Int

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150
func removeDuplicates(nums []int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}
