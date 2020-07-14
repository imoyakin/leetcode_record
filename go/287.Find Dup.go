package leetcode

//寻找重复数
func findDuplicate(nums []int) int {
	if len(nums) > 1 {
		slow := nums[0]
		fast := nums[nums[0]]
		for slow != fast {
			slow = nums[slow]
			fast = nums[nums[fast]]
		}
		entry := 0
		for entry != slow {
			entry = nums[entry]
			slow = nums[slow]
		}
		return entry
	}
	return -1
}
