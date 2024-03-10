package main

func containsDuplicate(nums []int) bool {
	kv := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := kv[v]; ok {
			return true
		}
		kv[v] = struct{}{}
	}
	return false
}
