package leetcode

//遍历
func searchInsert(nums []int, target int) int {
	for i := range nums {
		if nums[i] >= target {
			return i
		}
	}
	if nums[len(nums)-1] < target {
		return len(nums)
	}
	return 0
}

//二分
func searchInsert2(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	ans := n
	for left <= right {
		mid := (right-left)/2 + left
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
