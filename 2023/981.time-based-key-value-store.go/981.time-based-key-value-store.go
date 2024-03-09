package main

import "sort"

/*
 * @lc app=leetcode id=981 lang=golang
 *
 * [981] Time Based Key-Value Store
 *
 * https://leetcode.com/problems/time-based-key-value-store/description/
 *
 * algorithms
 * Medium (53.71%)
 * Likes:    3928
 * Dislikes: 377
 * Total Accepted:    309.8K
 * Total Submissions: 592.7K
 * Testcase Example:  '["TimeMap","set","get","get","set","get","get"]\n' +
  '[[],["foo","bar",1],["foo",1],["foo",3],["foo","bar2",4],["foo",4],["foo",5]]'
 *
 * Design a time-based key-value data structure that can store multiple values
 * for the same key at different time stamps and retrieve the key's value at a
 * certain timestamp.
 *
 * Implement the TimeMap class:
 *
 *
 * TimeMap() Initializes the object of the data structure.
 * void set(String key, String value, int timestamp) Stores the key key with
 * the value value at the given time timestamp.
 * String get(String key, int timestamp) Returns a value such that set was
 * called previously, with timestamp_prev <= timestamp. If there are multiple
 * such values, it returns the value associated with the largest
 * timestamp_prev. If there are no values, it returns "".
 *
 *
 *
 * Example 1:
 *
 *
 * Input
 * ["TimeMap", "set", "get", "get", "set", "get", "get"]
 * [[], ["foo", "bar", 1], ["foo", 1], ["foo", 3], ["foo", "bar2", 4], ["foo",
 * 4], ["foo", 5]]
 * Output
 * [null, null, "bar", "bar", null, "bar2", "bar2"]
 *
 * Explanation
 * TimeMap timeMap = new TimeMap();
 * timeMap.set("foo", "bar", 1);  // store the key "foo" and value "bar" along
 * with timestamp = 1.
 * timeMap.get("foo", 1);         // return "bar"
 * timeMap.get("foo", 3);         // return "bar", since there is no value
 * corresponding to foo at timestamp 3 and timestamp 2, then the only value is
 * at timestamp 1 is "bar".
 * timeMap.set("foo", "bar2", 4); // store the key "foo" and value "bar2" along
 * with timestamp = 4.
 * timeMap.get("foo", 4);         // return "bar2"
 * timeMap.get("foo", 5);         // return "bar2"
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= key.length, value.length <= 100
 * key and value consist of lowercase English letters and digits.
 * 1 <= timestamp <= 10^7
 * All the timestamps timestamp of set are strictly increasing.
 * At most 2 * 10^5 calls will be made to set and get.
 *
 *
*/

// @lc code=start

type values struct {
	vm       map[int]string
	sortedTs []int
}

type TimeMap struct {
	values map[string]values
}

func Constructor() TimeMap {
	return TimeMap{make(map[string]values)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	entry, ok := this.values[key]
	if !ok {
		entry = values{
			vm:       make(map[int]string),
			sortedTs: []int{},
		}
	}
	entry.vm[timestamp] = value
	prevTs := entry.sortedTs
	newTs := append(prevTs, timestamp)
	sort.Ints(newTs)
	entry.sortedTs = newTs
	this.values[key] = entry

}

func (this *TimeMap) Get(key string, timestamp int) string {
	allValues, ok := this.values[key]
	if !ok {
		return ""
	}
	timestamps := allValues.sortedTs
	first := timestamps[0]
	if timestamp < first {
		return ""
	}
	idx := binarySearch(timestamps, 0, len(timestamps)-1, timestamp)
	if idx == -1 {
		return ""
	}
	ts := timestamps[idx]
	return allValues.vm[ts]

}

func binarySearch(vs []int, l, r, v int) int {
	if r >= l {

		mid := l + (r-l)/2
		if v == vs[mid] {
			return mid
		}

		if vs[mid] > v {
			return binarySearch(vs, l, mid-1, v)
		}

		return binarySearch(vs, mid+1, r, v)
	}

	return r
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
// @lc code=end
