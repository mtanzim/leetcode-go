package main

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		v := nums[m]

		if v == target {
			return m
		}
		if target > v {
			l = m + 1
			continue
		}
		r = m - 1
	}
	return -1
}
