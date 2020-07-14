package leetcode

func twoSum(nums []int, target int) []int {
	var maps = make(map[int]int)
	for k, v := range nums {
		maps[v] = k
	}
	for i, j := range nums {
		c := target - j
		ret, ok := maps[c]
		if ret == i {
			continue
		}
		if ok {
			return []int{i, ret}
		}
	}
	return []int{0, 0}
}

func twoSum2(nums []int, target int) []int {
	//定一个map集合，然后将值为索引，健为值
	var maps = make(map[int]int)
	var arr []int
	//将值赋给map集合
	for k, i := range nums {
		maps[i] = k
	}

	//判断是否存在
	for k, i := range nums {
		curnt := target - i
		x1, x2 := maps[curnt]
		if x2 && x1 != k {
			arr = []int{x1, k}
		}
	}
	return arr
}

func twoSum3(nums []int, target int) []int {
	count := len(nums)
	var arr []int
	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			if nums[i]+nums[j] == target {
				arr = []int{i, j}
				break
			}
		}
	}

	return arr
}
