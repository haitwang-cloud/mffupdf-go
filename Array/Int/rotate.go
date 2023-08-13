package Int

func rotate(nums []int, k int) {
	k %= len(nums)
	swap(nums)
	swap(nums[:k])
	swap(nums[k:])

}

func swap(nums []int) {
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
