package Int

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
