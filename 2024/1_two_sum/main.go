package main

func twoSum(nums []int, target int) []int {
	kv := make(map[int]int)
	for ia, va := range nums {
		d := target - va
		if ib, ok := kv[va]; ok {
			return []int{ib, ia}
		}
		kv[d] = ia
	}
	return []int{}
}
